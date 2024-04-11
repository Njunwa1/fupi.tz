package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/ports"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/utils"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
	"log/slog"
	"os"
	"sync"
)

type payload struct {
	Request  *clicks.UrlClickRequest `json:"request"`
	Metadata metadata.MD             `json:"metadata"`
}

type Adapter struct {
	conn *amqp.Connection
	db   ports.DBPort
}

var (
	NumberOfWorkers = 5
)

func NewAdapter(rabbitUrl string, db ports.DBPort) *Adapter {
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		slog.Error("Cannot dial amqp", "err", err)
		os.Exit(1)
	}
	slog.Info("Connected to RabbitMQ", "url", rabbitUrl)
	return &Adapter{conn: conn, db: db}
}

func (a *Adapter) InitChannel() (*amqp.Channel, error) {
	return a.conn.Channel()
}

func (a *Adapter) ConsumeClickEvent() error {

	ch, err := a.InitChannel()
	if err != nil {
		slog.Error("Amq failed to init channel", "err", err)
		return err
	}

	q, err := ch.QueueDeclare(
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
	//buffered channel
	messageChannel := make(chan payload, NumberOfWorkers)

	for i := 1; i <= NumberOfWorkers; i++ {
		wg.Add(1) //counter++
		go a.worker(i, messageChannel, &wg)
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		false,  // AutoAck
		false,  // Exclusive
		false,  // NoLocal
		false,  // NoWait
		nil,    // Arguments
	)
	if err != nil {
		slog.Error("Failed to register a consumer:", "err", err)
		os.Exit(1)
	}

	// Start a goroutine to dispatch messages to workers
	//msgs in channel type that can only receive data
	go func() {
		for msg := range msgs {
			var message payload
			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				slog.Error("Error decoding message:", "error", err)
			} else {
				//Send message to buffered chan,
				//if the buffered chan is full this will block
				messageChannel <- message
				slog.Info("Received Click message for:", "shortUrl", message.Request.ShortUrl)
			}
			// Acknowledge the message
			_ = msg.Ack(false)
		}
		defer close(messageChannel) // Close the channel when no more messages are expected
	}()

	// Wait for all workers to finish
	wg.Wait()

	return nil
}

// worker is a goroutine that processes messages from the RabbitMQ queue
func (a *Adapter) worker(id int, messages <-chan payload, wg *sync.WaitGroup) {
	defer wg.Done() // --counter

	for message := range messages {
		slog.Info("Worker processing message:", "worker", id, "shortUrl", message.Request.ShortUrl)

		click, err := a.ProcessRequestMetadata(message.Metadata, message.Request)
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
	referer := utils.GetMD(md.Get("grpcgateway-referer"))
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

	urlIdObjectId, _ := primitive.ObjectIDFromHex(request.UrlID)
	click := domain.NewUrlClick(
		urlIdObjectId,
		userAgent,
		ipAdr,
		referer,
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
