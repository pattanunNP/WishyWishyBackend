package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pattanunNP/wishbackend/database"
	"github.com/pattanunNP/wishbackend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveWish(c *fiber.Ctx) error {

	req := new(models.User)
	err := c.BodyParser(req)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	creator_id := c.Locals("profile").(*models.LineProfile).UserID

	collection := database.MI.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": creator_id}

	change := bson.M{"$addToSet": bson.M{"wishs_list": req.WishID}}

	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(ctx, filter, change, opts)

	if err != nil {
		panic(err)
	}
	fmt.Println(&result)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
