package db

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (a *Adapter) SaveUser(ctx context.Context, user domain.User) error {
	collection := a.Client.Database("fupitz").Collection("users")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (a *Adapter) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	collection := a.Client.Database("fupitz").Collection("users")
	var result domain.User
	filter := domain.User{Email: email}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return domain.User{}, err
	}
	return result, nil
}
