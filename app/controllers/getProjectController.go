package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeabapps/project/app/databases"
	"github.com/jeabapps/project/app/models"
)

func GetProject(c *fiber.Ctx) error {
	db, err := databases.ConnectDB()
	if err != nil {
		return err
	}

	var project []models.Project
	if err := db.Find(&project).Error; err != nil {
		return err
	}

	return c.JSON(project)
}
