package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jeabapps/project/app/controllers"
)

func SetupApiRoutes(app *fiber.App, store *session.Store) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Azure Functions!")
	})

	// * Project, Unit
	api := app.Group("/api")

	// table
	// api.Get("/create-project-table", migrations.CreateProjectTable)
	// api.Get("/create-unit-table", migrations.CreateUintTable)

	// project
	api.Post("/create-project", controllers.CreateProject)
	api.Get("/get-project", controllers.GetProject)

	// unit
	api.Post("/create-unit", controllers.CreateUnit)
	api.Get("get-unit/:project_id", controllers.GetUnit)

}
