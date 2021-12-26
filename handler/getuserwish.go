package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pattanunNP/wishbackend/database"
	"github.com/pattanunNP/wishbackend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserWish(c *fiber.Ctx) error {

	user_id := c.Locals("profile").(*models.LineProfile).UserID

	db := database.MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	Pipeline := []bson.M{
		{"$match": bson.M{"user_id": user_id}},
		{"$unwind": bson.M{"path": "$wishs_list"}},
		{"$lookup": bson.M{
			"from":         "wishs",
			"localField":   "wishs_list",
			"foreignField": "wish_id",
			"as":           "userwish.userwish_info"}},
		{"$unwind": bson.M{"path": "$userwish.userwish_info"}},
		{"$group": bson.M{"_id": "$_id", "userwish_info": bson.M{"$push": "$userwish.userwish_info"}}},
	}

	cursor, err := db.Aggregate(ctx, Pipeline)
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	if err := cursor.Close(ctx); err != nil {
		panic(err)
	}

	fmt.Println(results)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    results,
	})
}
