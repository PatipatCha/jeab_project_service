package migrations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeabapps/project/app/databases"
	"github.com/jeabapps/project/app/models"
)

func CreateUintTable(c *fiber.Ctx) error {
	db, err := databases.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Unit{})

	return c.SendString("Successfully created Unit table")
}
