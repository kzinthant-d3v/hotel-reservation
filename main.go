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
	flag.Parse()

	var (
	app = fiber.New(config)
	port = flag.String("port", ":5000", "The port of api")
	userHandler = api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	hotelHandler = api.NewHotelHandler(hotelStore, roomStore)
	)

	apiv1 := app.Group("/api/v1")
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleListUsers)
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)

	app.Listen(*port)
}
