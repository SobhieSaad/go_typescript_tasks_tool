package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load("env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MONGO_DB_URI := os.Getenv("MONGO_DB_URI")

	clientOtions := options.Client().ApplyURI(MONGO_DB_URI)
	client, err := mongo.Connnect(context.Background(), clientOtions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mondo DB")

	collection = client.Database("todos").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getTodos(c *fiber.Ctx) error {
	return c.SendString("Get Todos")
}

func createTodo(c *fiber.Ctx) error {
	return c.SendString("Create Todo")
}

func updateTodo(c *fiber.Ctx) error {
	return c.SendString("Update Todo")
}

func deleteTodo(c *fiber.Ctx) error {
	return c.SendString("Delete Todo")
}
