package db

import (
	"context"
	"github.com/Njunwa1/fupi.tz/plan/internal/application/domain"
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

func (a *Adapter) Create(ctx context.Context, plans []domain.Plan) error {
	// Insert plans into the collection
	collection := a.Client.Database("fupitz").Collection("plans")
	for _, plan := range plans {
		_, err := collection.InsertOne(ctx, plan)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Adapter) Get(ctx context.Context, id string) (domain.Plan, error) {
	collection := a.Client.Database("fupitz").Collection("plans")
	var result domain.Plan
	filter := bson.M{"_id": id}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return domain.Plan{}, err
	}
	return result, nil
}

func (a *Adapter) GetAll(ctx context.Context) ([]domain.Plan, error) {
	collection := a.Client.Database("fupitz").Collection("plans")
	plansCursor, err := collection.Find(ctx, bson.D{})
	defer plansCursor.Close(ctx)
	var plans []domain.Plan
	if err = plansCursor.All(ctx, &plans); err != nil {
		return nil, err
	}
	return plans, nil
}
