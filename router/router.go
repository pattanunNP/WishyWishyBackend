package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pattanunNP/WishyWishyBackend/handler"
	"github.com/pattanunNP/WishyWishyBackend/middleware"
)

func SetupRoutes(app *fiber.App) {

	// Midelware
	api := app.Group("/api/v1")
	api.Get("/", handler.Hello)

	api.Get("/health", middleware.Authorization(), handler.Health)

	api.Post("/wish/createwish", middleware.Authorization(), handler.Createwish)

	api.Get("/wish/receivewish", middleware.Authorization(), handler.ReceiveWish)

	api.Get("/wish/getwish", middleware.Authorization(), handler.GetUserWish)

	api.Post("/user/savewish", middleware.Authorization(), handler.SaveWish)
}
