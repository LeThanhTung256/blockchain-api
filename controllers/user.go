package controllers

import (
	db "tayson/blockchain/database"
	"tayson/blockchain/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	db.DB.Find(&users)

	return c.JSON(users)
}
