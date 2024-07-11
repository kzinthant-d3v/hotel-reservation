package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kzinthant-d3v/hotel-reservation/types"
)

func HandleListUsers(c *fiber.Ctx) error{
	u := types.User{
		FirstName: "Kaskar",
		LastName: "Zin Thant",
	}
 return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error{
 return c.JSON(map[string]string{"user": "hello"})
}