package handler

import (
	"fmt"
	"log"

	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pattanunNP/wishbackend/database"
	"github.com/pattanunNP/wishbackend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserWishsList struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	WishID     string             `bson:"wish_id"`
	Content    string             `bson:"content"`
	Creator_id string             `bson:"creator_id"`
	Hidename   string             `bson:"hidename"`
	Created_at time.Time          `bson:"created_at"`
}

func GetUserWish(c *fiber.Ctx) error {

	user_id := c.Locals("profile").(*models.LineProfile).UserID

	filter := bson.M{"user_id": user_id}

	db := database.MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancel()

	
	matchStage := bson.D{{"$match", bson.D{{"user_id", user_id}}}}
	unwind_wishlist := bson.D{{"$unwind", bson.D{{"path", "$wishs_list"}, {"preserveNullAndEmptyArrays", false}}}}
	lookup := bson.D{{"$lookup", bson.D{{"from", "wishs"}, {"localField", "wishs_list"}, {"foreignField", "wish_id"}, {"as", "userwish.userwish_info"}}}}
	unwind_pack := bson.D{{"$unwind", bson.D{{"path", "$userwish.userwish_info"}, {"preserveNullAndEmptyArrays", false}}}}
	groupresult := bson.D{{"$group", bson.D{{"_id", "$_id"}} , {"userwish_info", bson.D{{"$push", "$userwish.userwish_info"}}}}}}
	
	cur,err := db.Aggregate(ctx, mongo.Pipeline{matchStage, unwind_wishlist, lookup, unwind_pack,groupresult})

	var Wishy []UserWishsList
	if err = cursor.All(ctx, &showsLoadedStruct); err != nil {
		panic(err)
	}
	fmt.Println(Wishy)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "result,"
	})
}
