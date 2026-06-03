package formregister

import (
	"strings"

	"github.com/axolotl-go/eternal_paw/internal/db"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Create(c *fiber.Ctx) error {
	var request Appointment

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Sanitización
	request.Name = strings.TrimSpace(request.Name)
	request.Phone = strings.TrimSpace(request.Phone)
	request.Service = strings.TrimSpace(request.Service)
	request.Species = strings.TrimSpace(request.Species)
	request.PetName = strings.TrimSpace(request.PetName)

	// Validación
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Guardar
	if err := db.DB.Create(&request).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Ok",
		"data":    request,
	})
}

func Views(c *fiber.Ctx) error {
	var requests []Appointment

	limit := c.QueryInt("limit", 50)
	page := c.QueryInt("page", 1)

	if limit > 100 {
		limit = 100
	}

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	if err := db.DB.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&requests).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	var total int64
	db.DB.Model(&Appointment{}).Count(&total)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": requests,
		"pagination": fiber.Map{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}
