package db

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/application/core/domain"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Adapter struct {
	Client *mongo.Client
}

func NewAdapter(dataSourceUrl string) *Adapter {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	//"mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dataSourceUrl)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
		return nil
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error while pinging DB: %s", err)
	}
	log.Println("Connected to MongoDB: ", dataSourceUrl)

	return &Adapter{Client: client}
}

func (a *Adapter) SaveUrlClick(ctx context.Context, click domain.UrlClick) error {
	collection := a.Client.Database("fupitz").Collection("clicks")
	_, err := collection.InsertOne(ctx, click)
	if err != nil {
		return err
	}
	return nil
}

func (a *Adapter) GetUrlClickAggregates(ctx context.Context, request *clicks.UserUrlRequest) (*clicks.UrlClicksAggregates, error) {
	collection := a.Client.Database("fupitz").Collection("clicks")
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var results []*clicks.UrlClicksAggregates
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	fmt.Println("Results: ", results)
	return results[0], nil
}
