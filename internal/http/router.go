package http

import (
	formregister "github.com/axolotl-go/eternal_paw/internal/form_register"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	api := app.Group("/api")

	api.Post("/form_register", formregister.Create)
	api.Get("/form_register", formregister.Views)

	// User
	// api.Get("/me", users.Me)
	// api.Post("/user", users.Create)
	// api.Post("/login", users.Login)
	// api.Post("/logout", users.Logout)
	// api.Post("/verify", users.Verify)

	// Pets
	// api.Post("/pets", pets.Create)
	// api.Get("/pet/:uuid", pets.Get)

	// Orders
	// api.Post("/orders", serviceorders.Create)

	// // SeriveType
	// api.Post("/service_type", servicetype.Create)
	// api.Delete("/service_type/:id", servicetype.Delete)
	// api.Get("/service_type", servicetype.Views)
	// api.Get("/service_type/:id", servicetype.View)

	// Certificacion
}
