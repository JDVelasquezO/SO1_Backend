package main

import (
	"backPractica1_SO1/Database"
	"backPractica1_SO1/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	if err := Database.Connect(); err != nil {
		log.Fatalln(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	Routes.Setup(app)
	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
