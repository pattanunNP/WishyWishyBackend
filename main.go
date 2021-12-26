package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pattanunNP/wishbackend/config"
	"github.com/pattanunNP/wishbackend/database"
	"github.com/pattanunNP/wishbackend/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.MongoDB()

	router.SetupRoutes(app)

	port := config.Config("APP_PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "4000"
	}
	fmt.Printf("Server started on port 🚀 %s\n", port)

	log.Fatal(app.Listen(":" + port))

}
