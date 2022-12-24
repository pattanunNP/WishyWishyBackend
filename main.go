package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pattanunNP/wishbackend/database"
	"github.com/pattanunNP/wishbackend/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.MongoDB()

	router.SetupRoutes(app)

	// If port is not provided, use default port

	if os.Getenv("PORT") != "" {
		port := os.Getenv("PORT")
		fmt.Printf("Server started on port 🚀 %s\n", port)

		log.Fatal(app.Listen(":"+port))
	
	}else{
		port:= 8080
		fmt.Printf("Server started on port 🚀 %d\n", port)

		log.Fatal(app.Listen(fmt.Sprintf(":%d",port)))
	
	}




}
