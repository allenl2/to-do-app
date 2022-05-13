package main

import (
	"log"
	"to-do-app/app/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi World! Welcome!")
	})

	app.Listen(":3000")

	database.Init()
	autoMigErr := database.AutoMigrateDB()

	if autoMigErr != nil {
		log.Println("Error occurred while auto migrating database")
	}
}
