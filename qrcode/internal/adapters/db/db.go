package db

import (
	"context"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/core/domain"
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

func (a *Adapter) Get(ctx context.Context, id string) (*domain.QRCode, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetAll(ctx context.Context) ([]*domain.QRCode, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) Update(ctx context.Context, qrCode *domain.QRCode) (*domain.QRCode, error) {
	//TODO implement me
	panic("implement me")
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

func (a *Adapter) Save(ctx context.Context, qrcode *domain.QRCode) (*domain.QRCode, error) {
	collection := a.Client.Database("fupitz").Collection("qrcodes")
	result, err := collection.InsertOne(ctx, qrcode)
	if err != nil {
		return &domain.QRCode{}, err
	}
	qrCodeId := result.InsertedID.(primitive.ObjectID)
	var qrCodeDoc domain.QRCode
	err = collection.FindOne(ctx, bson.M{"_id": qrCodeId}).Decode(&qrCodeDoc)
	if err != nil {
		return &domain.QRCode{}, err
	}
	return &qrCodeDoc, nil
}
