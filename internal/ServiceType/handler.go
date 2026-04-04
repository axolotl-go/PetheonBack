package servicetype

import (
	"github.com/axolotl-go/eternal_paw/internal/db"
	"github.com/axolotl-go/eternal_paw/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var serviceType ServiceType

	if err := c.BodyParser(&serviceType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if !utils.IsNotNull(serviceType.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Fields is required",
		})
	}

	if serviceType.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Price must be greater than 0",
		})
	}

	if err := db.DB.Where("name = ?", serviceType.Name).First(&serviceType).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Service type already exists",
		})
	}

	if err := db.DB.Create(&serviceType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create service type",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": serviceType,
	})
}

func View(c *fiber.Ctx) error {
	var serviceType ServiceType
	id := c.Params("id")

	if err := db.DB.Where("id = ?", id).First(&serviceType).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Service type not found",
		})
	}

	return c.JSON(fiber.Map{
		"name":        serviceType.Name,
		"description": serviceType.Description,
		"price":       serviceType.Price,
	})
}

func Views(c *fiber.Ctx) error {
	var serviceType []ServiceType

	if err := db.DB.Find(&serviceType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch service types",
		})
	}

	return c.JSON(serviceType)
}
