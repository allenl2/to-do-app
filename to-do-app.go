package main

import (
	"log"
	"to-do-app/app/controllers"
	"to-do-app/app/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Setup the server using Go FIber
	app := fiber.New()

	//Initialize the database
	database.Init()
	autoMigErr := database.AutoMigrateDB()

	if autoMigErr != nil {
		log.Println("Error occurred while auto migrating database")
	}

	//Routes
	//Gets
	//app.Get("/user/all", controllers.GetUsers)
	app.Get("/user", controllers.GetUser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi World! Welcome!")
	})

	app.Listen(":3000")
}
