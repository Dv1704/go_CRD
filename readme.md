
# 🚀 Simple Go MongoDB CRUD API

A robust and modern Go application demonstrating essential Create, Read, Update, and Delete (CRUD) operations for user management, backed by MongoDB. This project showcases a practical understanding of RESTful API design in Go, effective dependency management, and a successful migration from a deprecated database driver to the official, actively maintained solution.

## ✨ Features

This API provides the following core functionalities for managing user data:

* **Create User (POST /user)**: Add a new user to the database.
* **Get User by ID (GET /user/:id)**: Retrieve details of a specific user.
* **Delete User by ID (DELETE /user/:id)**: Remove a user record from the database.

## 🛠️ Tech Stack

* **Go (Golang)**: The primary programming language, known for its concurrency features and performance.
* **MongoDB**: A popular NoSQL document database, providing flexible data storage.
* **Go MongoDB Driver (go.mongodb.org/mongo-driver)**: The official, actively maintained Go driver for MongoDB.
* **HttpRouter (github.com/julienschmidt/httprouter)**: A lightweight and high-performance HTTP request router for Go.
* **Context (context)**: Utilized for managing request-scoped values, deadlines, and cancellation signals across API operations and database interactions.

## 🎯 The Challenge & The Solution: Driver Migration

Initially, this project utilized the `gopkg.in/mgo.v2` driver for MongoDB interactions. While mgo was once a popular choice, it is now deprecated and unmaintained, leading to potential compatibility issues with newer MongoDB server versions and a lack of ongoing support or security updates.

During development, this incompatibility manifested as a `Failed to connect to MongoDB: no reachable servers` error, despite the MongoDB server actively listening for connections. Upon deeper investigation via MongoDB server logs, it was revealed that connections were accepted but immediately terminated – a classic symptom of a protocol handshake mismatch due to the outdated driver.

My solution involved a complete migration to the official `go.mongodb.org/mongo-driver`. This refactoring addressed:

* **Dependency Management**: Updating go.mod and go.sum to reflect the new driver.
* **Connection Handling**: Rewriting the database session management using mongo.Connect and context.Context for robust connection establishment and timeouts.
* **CRUD Operations**: Adapting all database interaction logic within the controllers to the new driver's API, including FindOne, InsertOne, DeleteOne, and proper ObjectID handling using primitive.ObjectID.
* **Error Handling**: Implementing idiomatic Go error handling, logging, and appropriate HTTP status codes for improved API responsiveness and maintainability.

This migration not only resolved the immediate connection issue but also positioned the application for long-term stability, access to new MongoDB features, and adherence to modern Go development best practices.

## � Getting Started

Follow these steps to set up and run the application locally.

### Prerequisites

* **Go (1.22.4 or newer recommended)**: Ensure Go is installed and configured on your system.
* **MongoDB (4.0 or newer recommended)**: A running MongoDB instance.

### Installation

1. Clone the repository:

```bash
git clone https://github.com/dv1704/go_CRD.git 
cd go_CRD


2. Clean and Tidy Go Modules:
This command will fetch the necessary dependencies for the project, including the official MongoDB Go driver, and remove any deprecated ones.

```bash
go mod tidy
```

3. Run MongoDB

* Ensure your MongoDB server is running and accessible on `127.0.0.1:27017`.
* (e.g., `sudo systemctl start mongod` on Linux, or `brew services start mongodb-community` on macOS).
* Verify with `sudo ss -tulnp | grep ":27017"`.

## 🏃 Running the Application

Start the Go application:

```bash
go run main.go
```

You should see output similar to:

```
Database connected successfully!
Server is listening on port 9000...
```

### Access the API:

The API will be running on `http://localhost:9000`. You can use tools like curl, Postman, or your browser to interact with it.

## 💡 API Endpoints

Here are examples of how to interact with the API using curl:

### 1. Create a User (POST /user)

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "John Doe",
    "gender": "Male",
    "age": 30
}' http://localhost:9000/user
```

**Example Response:**

```json
{
    "id": "60c7b1a1f0a2c3d4e5f6a7b8",
    "name": "John Doe",
    "gender": "Male",
    "age": 30
}
```

### 2. Get User by ID (GET /user/:id)

Replace `60c7b1a1f0a2c3d4e5f6a7b8` with the actual ID returned from the Create User operation.

```bash
curl http://localhost:9000/user/60c7b1a1f0a2c3d4e5f6a7b8
```

**Example Response:**

```json
{
    "id": "60c7b1a1f0a2c3d4e5f6a7b8",
    "name": "John Doe",
    "gender": "Male",
    "age": 30
}
```

### 3. Delete User by ID (DELETE /user/:id)

Replace `60c7b1a1f0a2c3d4e5f6a7b8` with the actual ID of the user you want to delete.

```bash
curl -X DELETE http://localhost:9000/user/60c7b1a1f0a2c3d4e5f6a7b8
```

**Example Response:**

```
Deleted User 60c7b1a1f0a2c3d4e5f6a7b8
```

## 📂 Project Structure

```
simple_crud/
├── main.go                     # Entry point, sets up server and MongoDB connection
├── controllers/
│   └── user.go                 # Handles user-related API logic (CRUD operations)
├── models/
│   └── user.go                 # Defines the User struct and BSON/JSON tags
├── go.mod                      # Go module definition and dependencies
└── go.sum                      # Checksums for module dependencies
```

## 📈 Future Enhancements

* **Error Handling Refinement**: Implement custom error types and more granular error responses.
* **Input Validation**: Add more robust server-side validation for incoming request data.
* **Update Operation**: Implement an UPDATE endpoint for modifying user data.
* **Listing All Users**: Add a GET /users endpoint to retrieve a list of all users.
* **Testing**: Implement unit and integration tests for controllers and database interactions.
* **Dockerization**: Containerize the application and MongoDB for easier deployment.
* **Configuration Management**: Externalize database connection strings and other settings.

## 🤝 Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests.

## 📄 License

This project is licensed under the MIT License - see the LICENSE.md file for details.
