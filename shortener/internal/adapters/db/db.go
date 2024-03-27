package db

import (
	"context"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"log"
	"time"
)

type Adapter struct {
	Client *mongo.Client
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	//"mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(dataSourceUrl)
	clientOptions.Monitor = otelmongo.NewMonitor()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
		return nil, err
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error while pinging DB: %s", err)
	}
	log.Println("Connected to MongoDB: ", dataSourceUrl)

	return &Adapter{Client: client}, nil
}

func (a *Adapter) SaveUrl(ctx context.Context, url domain.Url) (*domain.Url, error) {
	collection := a.Client.Database("fupitz").Collection("urls")
	result, err := collection.InsertOne(ctx, url)
	if err != nil {
		return &domain.Url{}, err
	}
	urlID := result.InsertedID.(primitive.ObjectID)
	var urlDoc domain.Url
	err = collection.FindOne(ctx, bson.M{"_id": urlID}).Decode(&urlDoc)
	if err != nil {
		return &domain.Url{}, err
	}
	return &urlDoc, nil
}

func (a *Adapter) GetUrlByShortUrl(ctx context.Context, shortUrl string) (domain.Url, error) {
	collection := a.Client.Database("fupitz").Collection("urls")
	var result domain.Url
	filter := bson.D{{"short", shortUrl}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return domain.Url{}, err
	}
	return result, nil
}

func (a *Adapter) GetAllUserUrls(ctx context.Context, userId *primitive.ObjectID) ([]domain.Url, error) {
	collection := a.Client.Database("fupitz").Collection("urls")
	filter := bson.M{"user_id": userId}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var urls []domain.Url
	if err = cursor.All(ctx, &urls); err != nil {
		return nil, err
	}
	return urls, nil
}
