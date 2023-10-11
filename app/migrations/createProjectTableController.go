package migrations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeabapps/project/app/databases"
	"github.com/jeabapps/project/app/models"
)

func CreateProjectTable(c *fiber.Ctx) error {
	db, err := databases.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Project{})

	return c.SendString("Successfully created Project table")
}
