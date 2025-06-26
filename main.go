package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"simple_crud/controllers"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// This function now returns a *mongo.Client and an error
	client, err := getMongoClient()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		// Ensure the client is disconnected when main exits
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB.")
	}()


	uc := controllers.NewUserController(client)

	// Initialize the httprouter
	r := httprouter.New()

	// Define routes
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// Start the server
	fmt.Println("Server is listening on port 9000...")
	serverErr := http.ListenAndServe(":9000", r)
	if serverErr != nil {
		log.Fatalf("Server failed to start: %v", serverErr) 
	}
}

// getMongoClient establishes a connection to MongoDB using the official driver.
func getMongoClient() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Create a context with a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the primary to verify connection
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer pingCancel()
	err = client.Ping(pingCtx, nil)
	if err != nil {
		// Log the underlying reason for ping failure
		log.Printf("MongoDB Ping failed: %v", err)
		// Disconnect immediately if ping fails
		if disconnectErr := client.Disconnect(context.TODO()); disconnectErr != nil {
			log.Printf("Error during disconnect after failed ping: %v", disconnectErr)
		}
		return nil, fmt.Errorf("MongoDB server not reachable or responsive: %w", err)
	}

	fmt.Println("Database connected successfully!")
	return client, nil
}
