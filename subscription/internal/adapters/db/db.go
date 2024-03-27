package db

import (
	"context"
	"github.com/Njunwa1/fupi.tz/subscription/internal/application/domain"
	"go.mongodb.org/mongo-driver/bson"
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

func (a *Adapter) CreateSubscription(ctx context.Context, subscription domain.Subscription) (domain.Subscription, error) {
	collection := a.Client.Database("fupitz").Collection("subscriptions")
	res, err := collection.InsertOne(ctx, subscription)
	if err != nil {
		return domain.Subscription{}, err
	}
	var result domain.Subscription
	err = collection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&result)
	if err != nil {
		return domain.Subscription{}, err
	}
	return result, nil
}

func (a *Adapter) GetUserActiveSubscription(ctx context.Context, userId string) (domain.Subscription, error) {
	collection := a.Client.Database("fupitz").Collection("subscriptions")
	var result domain.Subscription
	filter := bson.M{"user_id": userId, "is_active": true}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return domain.Subscription{}, err
	}
	return result, nil
}

func (a *Adapter) GetAll(ctx context.Context) ([]domain.Subscription, error) {
	collection := a.Client.Database("fupitz").Collection("subscriptions")
	plansCursor, err := collection.Find(ctx, bson.D{})
	defer plansCursor.Close(ctx)
	var plans []domain.Subscription
	if err = plansCursor.All(ctx, &plans); err != nil {
		return nil, err
	}
	return plans, nil
}
