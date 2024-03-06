package db

import (
	"context"
	"github.com/njunwa1/fupi.tz/keygen/internal/application/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Adapter struct {
	client *mongo.Client
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	//"mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dataSourceUrl)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return &Adapter{client: client}, nil
}

// SaveShortUrlKey implicitly implements the DBPort interface
func (a *Adapter) SaveShortUrlKey(ctx context.Context, entry domain.KeyGenLogEntry) error {
	collection := a.client.Database("fupi.tz").Collection("keygen")
	_, err := collection.InsertOne(ctx, entry)
	if err != nil {
		return err
	}
	return nil
}
