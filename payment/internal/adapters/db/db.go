package db

import (
	"context"
	"github.com/Njunwa1/fupi.tz/payment/internal/application/domain"
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
	clientOptions := options.Client()
	clientOptions.ApplyURI(dataSourceUrl)
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

func (a *Adapter) CreatePayment(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	collection := a.Client.Database("fupitz").Collection("payments")
	insertedResult, err := collection.InsertOne(ctx, payment)
	if err != nil {
		return domain.Payment{}, err
	}
	var result domain.Payment
	err = collection.FindOne(ctx, bson.M{"_id": insertedResult.InsertedID}).Decode(&result)
	if err != nil {
		return domain.Payment{}, err
	}
	return result, nil
}
func (a *Adapter) GetUserPayments(ctx context.Context, userId string) ([]domain.Payment, error) {
	collection := a.Client.Database("fupitz").Collection("payments")
	cursor, err := collection.Find(ctx, bson.M{"user_id": userId})
	defer cursor.Close(ctx)
	var results []domain.Payment
	if err != nil {
		return []domain.Payment{}, err
	}
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (a *Adapter) GetAllPayments(ctx context.Context) ([]domain.Payment, error) {
	collection := a.Client.Database("fupitz").Collection("payments")
	cursor, err := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	var results []domain.Payment
	if err != nil {
		return []domain.Payment{}, err
	}
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
