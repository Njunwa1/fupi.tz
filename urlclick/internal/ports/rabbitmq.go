package ports

type RabbitMQPort interface {
	// ConsumeClickEvent is a method that is called when a user clicks on a short url
	ConsumeClickEvent() error
}
