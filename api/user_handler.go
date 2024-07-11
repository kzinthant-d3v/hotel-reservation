package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kzinthant-d3v/hotel-reservation/db"
	"github.com/kzinthant-d3v/hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler (userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error{
	var (
	id = c.Params("id");
	ctx = context.Background()
	)

	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
 return c.JSON(user)
}
func (h *UserHandler) HandleListUsers(c *fiber.Ctx) error{
	u := types.User{
		FirstName: "Kaskar",
		LastName: "Zin Thant",
	}
 return c.JSON(u)
}
