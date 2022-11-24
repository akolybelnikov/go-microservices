package main

import (
	"context"
	"github.com/akolybelnikov/go-microservices/logger-service/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

const (
	webPort  = "80"
	mongoURL = "mongodb://mongo:27017"
)

var client *mongo.Client

type Config struct {
	models data.Models
}

func main() {
	// connect to MongoDB
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

	client = mongoClient

	// create disconnect context
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// disconnect from MongoDB
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Could not disconnect from MongoDB: %v", err)
		}
	}()

	app := &Config{
		models: data.New(client),
	}

	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	log.Println("Server started on port", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return c, nil
}
