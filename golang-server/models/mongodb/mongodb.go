package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGODB_URI string = os.Getenv("MONGODB_URI")

var client *mongo.Client

// Connect initializes the MongoDB client and connects to the server
func Connect(uri string) {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

// GetClient returns the MongoDB client
func GetClient() *mongo.Client {
	return client
}

func GeColl(coll string) *mongo.Collection {
	if client == nil {
		log.Panic("MongoDB client does not exists")
	}
	return client.Database("pmp").Collection(coll)
}

// Disconnect disconnects the MongoDB client
func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
func init() {
	log.Println("MONGODB_URI ", MONGODB_URI)
	// Connect(os.Getenv(MONGODB_URI))
	Connect("mongodb://localhost:27017/")
}
