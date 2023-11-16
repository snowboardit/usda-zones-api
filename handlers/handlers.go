package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowboardit/usda-zones-api/lib/data"
	"github.com/snowboardit/usda-zones-api/lib/zone"
)

type Row = data.Row

var store = data.Store

func GetByZip(c *fiber.Ctx) error {
	code := c.Params("code")
	res, err := zone.GetZoneByZip(code, store)
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
