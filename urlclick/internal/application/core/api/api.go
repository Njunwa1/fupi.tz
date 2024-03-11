package api

import (
	"github.com/Njunwa1/fupi.tz/urlclick/internal/ports"
)

type Application struct {
	db       ports.DBPort
	rabbitmq ports.RabbitMQPort
}

func NewApplication(db ports.DBPort, rabbitmq ports.RabbitMQPort) *Application {
	return &Application{db: db, rabbitmq: rabbitmq}
}
