package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pattanunNP/WishyWishyBackend/database"
	"github.com/pattanunNP/WishyWishyBackend/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.MongoDB()

	router.SetupRoutes(app)

	// If port is not provided, use default port

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
	}
	fmt.Printf("Server started on port 🚀 %s\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

}
