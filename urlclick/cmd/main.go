package cmd

import (
	"context"
	"github.com/Njunwa1/fupi.tz/urlclick/config"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/adapters/rabbitmq"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func main() {
	dbAdapter := db.NewAdapter(config.GetDataSourceURL())
	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			slog.Error("Failed to disconnect from the database", "err", err)
		}
	}(dbAdapter.Client, context.Background())

	rabbitmqAdapter := rabbitmq.NewAdapter(config.GetRabbitMQURL(), dbAdapter)
	err := rabbitmqAdapter.ConsumeClickEvent()
	if err != nil {
		slog.Error("Failed to consume click event", "err", err)
	}

}
