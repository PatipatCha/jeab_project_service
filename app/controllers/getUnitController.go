package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeabapps/project/app/databases"
	"github.com/jeabapps/project/app/models"
)

func GetUnit(c *fiber.Ctx) error {
	projectID := c.Params("project_id")

	db, err := databases.ConnectDB()
	if err != nil {
		return err
	}

	var unit []models.Unit
	if err := db.Where("project_id = ?", projectID).Find(&unit).Error; err != nil {
		return err
	}

	return c.JSON(unit)
}
