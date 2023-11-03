package database

import (
	"context"
	"log"
	"time"

	"github.com/pattanunNP/WishyWishyBackend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func MongoDB() {
	// Add your init code here
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config("MONGODB_URI")))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	db := client.Database("production")

	MI = MongoInstance{
		Client: client,
		DB:     db,
	}

}
