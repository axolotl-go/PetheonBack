package serviceorders

import (
	"github.com/axolotl-go/eternal_paw/internal/db"
	"github.com/axolotl-go/eternal_paw/internal/pets"
	"github.com/axolotl-go/eternal_paw/internal/users"
	"github.com/axolotl-go/eternal_paw/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var order Order
	var user users.User
	var pet pets.Pet

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := db.DB.Where("id = ?", order.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if err := db.DB.Where("id = ?", order.PetID).First(&pet).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Pet not found",
		})
	}

	if pet.UserID != user.ID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Pet does not belong to user",
		})
	}

	order.ServiceType = "individual"
	order.Status = "pending"

	if order.PickupRequired {
		order.PickupAddress = user.Address
	} else {
		order.PickupAddress = ""
	}

	// Urn

	var pricing float64

	if order.ServiceType == "composta" {
		pricing = (pet.Weight * 90) + 500
	} else if order.ServiceType == "cremation" {
		pricing = (pet.Weight * 120) + 500
	} else {
		pricing = pet.Weight * 120
	}
	order.Price = pricing

	order.OrderNumber = utils.GenerateOrder()

	if err := db.DB.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create order",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"order": order,
	})
}
