package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/ports"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc/metadata"
	"log"
	"log/slog"
	"sync"
)

type payload struct {
	request  *clicks.UrlClickRequest
	metadata metadata.MD
}

type Adapter struct {
	conn *amqp.Connection
	db   ports.DBPort
}

func NewAdapter(rabbitUrl string, db ports.DBPort) *Adapter {
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		log.Printf("Cannot dial amqp %s", err)
	}
	return &Adapter{conn: conn, db: db}
}

func (a *Adapter) initChannel() (*amqp.Channel, error) {
	return a.conn.Channel()
}

func (a *Adapter) ConsumeClickEvent() error {

	ch, err := a.initChannel()
	if err != nil {
		slog.Error("Amq failed to init channel", "err", err)
		return err
	}

	_, err = ch.QueueDeclare(
		"url_clicked", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	if err != nil {
		slog.Error("Amq failed to declare Queue", "err", err)
		return err
	}

	// Create a worker pool
	var wg sync.WaitGroup
	messageChannel := make(chan payload, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go a.worker(i, messageChannel, &wg)
	}

	// Consume messages
	msgs, err := ch.Consume(
		"url_clicked", // Queue name
		"",            // Consumer tag
		false,         // AutoAck
		false,         // Exclusive
		false,         // NoLocal
		false,         // NoWait
		nil,           // Arguments
	)
	if err != nil {
		log.Fatal("Failed to register a consumer:", err)
	}

	// Start a goroutine to dispatch messages to workers
	go func() {
		for msg := range msgs {
			var message payload
			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				slog.Error("Error decoding message:", "error", err)
				// Optionally handle the decoding error
			} else {
				messageChannel <- message
			}
			// Acknowledge the message
			msg.Ack(false)
		}
		close(messageChannel) // Close the channel when no more messages are expected
	}()

	// Wait for all workers to finish
	wg.Wait()

	return nil
}

// worker is a goroutine that processes messages from the RabbitMQ queue
func (a *Adapter) worker(id int, messages <-chan payload, wg *sync.WaitGroup) {
	defer wg.Done()

	for message := range messages {
		log.Printf("Worker %d processing message: %s %s", id, message.request, message.metadata)

		click, err := a.ProcessRequestMetadata(message.metadata, message.request)
		if err != nil {
			slog.Error("Error processing message: ", "err", err)
		}
		err = a.db.SaveUrlClick(context.Background(), click)
		if err != nil {
			slog.Error("Error saving message to the database: ", "err", err)
		}
	}
}

func (a *Adapter) ProcessRequestMetadata(md metadata.MD, request *clicks.UrlClickRequest) (domain.UrlClick, error) {
	userAgent := utils.GetMD(md.Get("grpcgateway-user-agent"))
	ipAdr := utils.GetMD(md.Get("x-forwarded-for"))

	//TODO replace hard-code "" with ipAdr
	var country string
	var city string
	var longitude float64
	var latitude float64
	record, err := utils.GetIPInfo("197.186.19.95")
	if err != nil {
		slog.Error("error getting ip info: %v", err)
		return domain.UrlClick{}, err
	} else {
		city = record.City.Names["en"]
		country = record.Country.Names["en"]
		longitude = record.Location.Longitude
		latitude = record.Location.Latitude
	}

	click := domain.NewUrlClick(
		request.UrlID,
		userAgent,
		ipAdr,
		"",
		utils.DeviceFromUserAgent(userAgent),
		utils.BrowserFromUserAgent(userAgent),
		utils.OSFromUserAgent(userAgent),
		country,
		city,
		latitude,
		longitude,
	)

	return click, nil
}
