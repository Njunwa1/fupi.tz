package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc/metadata"
	"log"
	"log/slog"
	"os"
)

type Adapter struct {
	conn *amqp.Connection
}

type payload struct {
	Request  *clicks.UrlClickRequest `json:"request"`
	Metadata metadata.MD             `json:"metadata"`
}

func NewAdapter(rabbitUrl string) *Adapter {
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		log.Printf("Cannot dial amqp %s", err)
		os.Exit(1)
	}
	slog.Info("Connected to RabbitMQ", "url", rabbitUrl)
	return &Adapter{conn: conn}
}

func (a *Adapter) initChannel() (*amqp.Channel, error) {
	return a.conn.Channel()
}

func (a *Adapter) PublishClickEvent(ctx context.Context, request *clicks.UrlClickRequest, md metadata.MD) error {

	ch, err := a.initChannel()
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

	body := payload{
		Request:  request,
		Metadata: md,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		slog.Error("Failed to marshal JSON: %v", err)
		return err
	}
	slog.Info("Publishing JSON Data to RabbitMQ", "data", jsonData)

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})
	return nil
}
