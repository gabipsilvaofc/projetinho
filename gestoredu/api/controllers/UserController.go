package controllers

import (
	"net/http"
	
	"teste/api/entities"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	users []entities.User
}

func NewUserController() *userController {
	return &userController{}
}

//delete
func (t *userController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	for idx, user := range t.users {
		if user.ID == id {
			t.users = append(t.users[0:idx], t.users[idx+1:]...)
			return ctx.SendStatus(http.StatusOK)
			}
		}
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
	})
}