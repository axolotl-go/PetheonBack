package http

import (
	servicetype "github.com/axolotl-go/eternal_paw/internal/ServiceType"
	"github.com/axolotl-go/eternal_paw/internal/pets"
	serviceorders "github.com/axolotl-go/eternal_paw/internal/service_orders"
	"github.com/axolotl-go/eternal_paw/internal/users"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	api := app.Group("/api")

	// User
	api.Post("/user", users.Create)
	api.Get("/user", users.GetUserData)
	api.Post("/login", users.Login)
	api.Post("/logout", users.Logout)
	api.Post("/verify", users.Verify)

	// Pets
	api.Post("/pets", pets.Create)
	api.Get("/pet/:uuid", pets.Get)

	// Orders
	api.Post("/orders", serviceorders.Create)

	// SeriveType
	api.Post("/service_type", servicetype.Create)
	api.Get("/service_type", servicetype.Views)
	api.Get("/service_type/:id", servicetype.View)

	// Certificatiiobn
}
