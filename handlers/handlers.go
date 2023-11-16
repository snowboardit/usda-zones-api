package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/snowboardit/usda-zones-api/lib/data"
	"github.com/snowboardit/usda-zones-api/lib/zone"
)

type Row = data.Row

var Data []Row

func GetByZip(c *fiber.Ctx) error {
	// If we don't have data, load it
	if Data == nil {
		fmt.Println("🕦 Loading data...")
		Data = data.Load()
		if Data == nil {
			panic("❌ Unable to load data")
		}
		fmt.Println("✅ Data loaded")
	}
	code := c.Params("code")
	res, err := zone.GetZoneByZip(code, Data)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":      false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "",
		"data":    res,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}
