package main

import (
	"log"

	servicetype "github.com/axolotl-go/eternal_paw/internal/ServiceType"
	"github.com/axolotl-go/eternal_paw/internal/config"
	"github.com/axolotl-go/eternal_paw/internal/db"
	"github.com/axolotl-go/eternal_paw/internal/http"
	"github.com/axolotl-go/eternal_paw/internal/pets"
	serviceorders "github.com/axolotl-go/eternal_paw/internal/service_orders"
	"github.com/axolotl-go/eternal_paw/internal/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var serverxport string

func init() {
	cfg := config.Load()
	serverxport = cfg.ServerPort

	if serverxport == "" {
		serverxport = "3000"
	}
}

func main() {

	db.DB.AutoMigrate(
		&users.User{},
		&pets.Pet{},
		&serviceorders.Order{},
		&servicetype.ServiceType{},
	)

	app := fiber.New()
	app.Use(cors.New(config.CorsConfig()))
	app.Static("/storage", "/var/app/storage")
	app.Static("/", "../../index.html")

	http.SetupRouter(app)

	log.Fatal(app.Listen(":" + serverxport))
}
