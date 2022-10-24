package main

import (
	"User_Management/Database"
	"User_Management/Routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	MONGODB_URL = "mongodb://localhost:27017/users_db"
)

func main() {
	app := fiber.New()

	err := Database.ConnectDB(MONGODB_URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	Routes.UsermManagement(app)
	app.Listen(":6000")
}
