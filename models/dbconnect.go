package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDB used to connect MongoDB.
func ConnectToDB() *mongo.Client {
	// Construct MongoDB connection URI.
	mongoURI := fmt.Sprintf("%s://%s:%s", os.Getenv("DB_DRIVER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	// Context with timeout, if connecting to the database takes longer
	// then 10 seconds, cancel the connection.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the database, if there is an error, log the error
	// and exit the program.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	return client
}
