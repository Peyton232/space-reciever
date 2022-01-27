package database

import (
	"context"
	"log"
	"time"

	model "github.com/Peyton232/dual-go-reciever/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client   *mongo.Client
	dataIrid *mongo.Collection
}

func Connect() *DB {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ubxcq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Print(err)
		log.Print("\nDB connection failed in database package")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	return &DB{
		dataIrid: client.Database("nft-space").Collection("dataIrid"),
		client:   client,
	}
}

//Creation
func (db *DB) SendData(input *model.NewSpaceData) *model.SpaceData {
	collection := db.dataIrid
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newID := time.Now()
	data := &model.SpaceData{
		SpaceID:          newID.String(),
		Imei:             input.Imei,
		Serial:           input.Serial,
		Momsn:            input.Momsn,
		TransmitTime:     input.TransmitTime,
		IridiumLatitude:  input.IridiumLatitude,
		IridiumLongitude: input.IridiumLongitude,
		IridiumCep:       input.IridiumCep,
		Data:             input.Data,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil || res == nil {
		log.Print(err)
		log.Print("\nunable to insert user into DB in database package\n")
		return nil
	}
	return data
}

// Disconnect to MongoDB - could use, on further inspection looks like mongo is good at closing itself
func (db *DB) Disconnect() {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Print(err)
		log.Print("\nunable to disconnect from DB in database package\n")
	}
}
