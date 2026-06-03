package pets

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/axolotl-go/eternal_paw/internal/config"
	"github.com/axolotl-go/eternal_paw/internal/db"
	"github.com/axolotl-go/eternal_paw/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var pet Pet

	pet.Name = c.FormValue("name")
	pet.Breed = c.FormValue("breed")
	pet.Species = c.FormValue("species")

	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil || userID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user_id",
		})
	}
	// pet.UserID = uint(userID)

	weight, err := strconv.ParseFloat(c.FormValue("weight"), 64)
	if err != nil || weight <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid weight",
		})
	}
	pet.Weight = weight

	deathDateStr := c.FormValue("death_date")
	if deathDateStr != "" {
		t, err := time.Parse("2006-01-02", deathDateStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid death_date format (YYYY-MM-DD required)",
			})
		}
		pet.DeathDate = &t
	}

	if pet.Name == "" || pet.Breed == "" || pet.Species == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	file, err := c.FormFile("image")
	if err == nil && file != nil {

		ext := strings.ToLower(filepath.Ext(file.Filename))

		allowed := map[string]bool{
			".png":  true,
			".jpg":  true,
			".jpeg": true,
		}

		if !allowed[ext] {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid file type",
			})
		}

		filename := utils.UUIDGenerator(pet.Name) + ext
		imgPath := config.Load().StoragePath + "/images/" + filename

		if err := c.SaveFile(file, imgPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error saving image",
			})
		}

		pet.ImageUrl = &filename
	}

	if err := db.DB.Create(&pet).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating pet",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Pet created successfully",
		"pet":     pet,
	})
}

func Get(c *fiber.Ctx) error {
	params := c.Params("uuid")

	path := config.Load().StoragePath

	return c.SendFile(path + "/images/" + params)
}
