package handler

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pattanunNP/WishyWishyBackend/database"

	"github.com/pattanunNP/WishyWishyBackend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReceiveWish(c *fiber.Ctx) error {
	rand.Seed(time.Now().Unix())
	db := database.MI.DB.Collection("wishs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := db.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	defer cur.Close(ctx)

	var wishs []models.Wishy
	for cur.Next(ctx) {
		var wish models.Wishy
		err := cur.Decode(&wish)
		if err != nil {
			fmt.Println(err)
		}
		wishs = append(wishs, wish)
	}

	n := rand.Int() % len(wishs)

	fmt.Println(n)

	wish := wishs[n]
	fmt.Printf("%+v\n", wish)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    wish,
	})
}
