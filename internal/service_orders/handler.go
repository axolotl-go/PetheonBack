package serviceorders

import (
	servicetype "github.com/axolotl-go/eternal_paw/internal/ServiceType"
	"github.com/axolotl-go/eternal_paw/internal/db"
	"github.com/axolotl-go/eternal_paw/internal/users"
	"github.com/axolotl-go/eternal_paw/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Create(c *fiber.Ctx) error {
	var order Order
	var user users.User
	var serviceType servicetype.ServiceType

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := db.DB.Where("id = ?", order.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if err := db.DB.Where("id = ?", order.ServiceTypeID).First(&serviceType).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Service not found",
		})
	}

	order.Active = false
	order.OrderNumber = utils.GenerateOrder()

	if err := validate.Struct(order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Pricing

	// Address
	order.Status = "pending"

	order.PickupAddress = ""
	if order.PickupRequired && user.Address != "" {
		order.PickupAddress = user.Address
	}

	// if err := db.DB.Create(&order).Error; err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Created successfully",
		"Order":   order,
	})
}

func Views(c *fiber.Ctx) error {
	return c.JSON("Good")
}
