package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeabapps/project/app/databases"
	"github.com/jeabapps/project/app/models"
)

func CreateUnit(c *fiber.Ctx) error {
	var request models.Unit
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	request.Status = "active"

	db, err := databases.ConnectDB()
	if err != nil {
		return err
	}

	if err := db.Create(&request).Error; err != nil {
		return err
	}

	return c.SendString("Success Created Unit")
}
