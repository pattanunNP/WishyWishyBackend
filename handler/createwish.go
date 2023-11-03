package handler

import (
	"fmt"
	"strings"
	"time"

	"context"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/pattanunNP/WishyWishyBackend/database"
	"github.com/pattanunNP/WishyWishyBackend/models"
)

func Createwish(c *fiber.Ctx) error {

	// Get the request body

	req := new(models.Wish)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	wish := models.Wish{

		Content:     req.Content,
		Creator:     req.Creator,
		Hidename:    req.Hidename,
		CreatedTime: time.Now(),
	}

	// fmt.Printf("%+v\n", wish)
	_id := uuid.New().String()
	wish_id := strings.Replace(_id, "-", "", -1)
	// fmt.Println(wish_id)
	collection := database.MI.DB.Collection("wishs")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	creator_id := c.Locals("profile").(*models.LineProfile).UserID

	result, err := collection.InsertOne(ctx, &models.Wishy{
		WishID:      wish_id,
		Content:     wish.Content,
		Creator:     wish.Creator,
		CreaterId:   creator_id,
		Hidename:    wish.Hidename,
		CreatedTime: wish.CreatedTime,
	})
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Createdwish",
	})
}
