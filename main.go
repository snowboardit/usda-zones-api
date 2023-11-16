package main

import (
	"flag"
	"log"

	"github.com/snowboardit/usda-zones-api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in production")
)

func main() {

	// Parse cmd line flags
	flag.Parse()

	// Create new Fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Setup middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create api group
	api := app.Group("/api")

	api.Get("/zip/:code", handlers.GetByZip)

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
