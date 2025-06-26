package controllers

import (
	"context" 
	"encoding/json"
	"fmt"
	"log" 
	"net/http"
	"time" 

	"simple_crud/models" 

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson" 
	"go.mongodb.org/mongo-driver/bson/primitive" 
	"go.mongodb.org/mongo-driver/mongo"         
)

// UserController holds the MongoDB Collection for users
type UserController struct {
	usersCollection *mongo.Collection 
}

// NewUserController creates a new UserController instance.
func NewUserController(client *mongo.Client) *UserController {
	databaseName := "Mongo_golang"
	collectionName := "User"
	return &UserController{
		usersCollection: client.Database(databaseName).Collection(collectionName),
	}
}

// GetUser handles fetching a single user by ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idStr := p.ByName("id")

	// Validate and parse the ID string into a MongoDB ObjectID
	oid, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest) 
		return
	}

	u := models.User{}
	// Create a context with a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find the user by ID
	filter := bson.M{"_id": oid} 
	err = uc.usersCollection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		log.Printf("Error fetching user: %v", err) 
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Marshal the user object to JSON
	uj, err := json.Marshal(u)
	if err != nil {
		log.Printf("Error marshaling user to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") 
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(uj) // Use w.Write for byte slices
	if err != nil {
		log.Printf("Error writing response: %v", err) 
	}
}

// CreateUser handles creating a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	// Decode the request body into the User struct
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a new ObjectID for the user
	u.ID = primitive.NewObjectID()

	// Create a context with a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert the user into the collection
	_, err = uc.usersCollection.InsertOne(ctx, u)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Internal server error during user creation", http.StatusInternalServerError)
		return
	}

	// Marshal the created user object to JSON (including its new ID)
	uj, err := json.Marshal(u)
	if err != nil {
		log.Printf("Error marshaling new user to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	_, err = w.Write(uj)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// DeleteUser handles deleting a user by ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idStr := p.ByName("id")

	// Validate and parse the ID string into a MongoDB ObjectID
	oid, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	// Create a context with a timeout for the database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the user by ID
	filter := bson.M{"_id": oid}
	res, err := uc.usersCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Internal server error during user deletion", http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain") 
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Deleted User %s\n", oid.Hex()) 
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}