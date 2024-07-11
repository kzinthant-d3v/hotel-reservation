package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kzinthant-d3v/hotel-reservation/api"
	"github.com/kzinthant-d3v/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation-db"
const userCollection = "users"

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

  client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	port := flag.String("port", ":5000", "The port of api");
	flag.Parse()

  app := fiber.New(config)

	 //handler initialization 
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client));

	apiv1 := app.Group("/api/v1")
	apiv1.Get("/foo", handleFoo)
	apiv1.Get("/user", userHandler.HandleListUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*port)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "hello foo"})
}