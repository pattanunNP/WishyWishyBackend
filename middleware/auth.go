package middleware

import (
	"strings"

	"github.com/pattanunNP/wishbackend/util"

	"github.com/gofiber/fiber/v2"
)

func Authorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		token := strings.Split(authorization, "Bearer")
		// fmt.Println(token)

		if len(token[1]) <= 1 {
			return c.Status(401).JSON(fiber.Map{"message": "Missing or malformed JWT"})
		} else {

			profile, err := util.Decode(token[1])
			if err != nil {
				return c.Status(403).JSON(fiber.Map{"message": err.Error()})
			}
			c.Locals("profile", profile)

			return c.Next()

		}

	}

}
