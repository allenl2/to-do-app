package main

import (
	"log"
	"to-do-app/app/controllers"
	"to-do-app/app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	//Setup the server using Go Fiber
	app := fiber.New()

	//Initialize the database
	database.Init()
	autoMigErr := database.AutoMigrateDB()

	if autoMigErr != nil {
		log.Println("Error occurred while auto migrating database")
	}

	//Initialize redis cache and session
	database.InitRedis()

	//Authenication middelware
	app.Use("/tasks", basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "123456",
		},
		Realm:      "Forbidden",
		Authorizer: controllers.LoginAuth,
	}))

	//ROUTES

	//Users (Temporary) - will replace with Auth service
	app.Get("/user/:username", controllers.GetUser)
	app.Post("/user", controllers.CreateUser)

	//Tasks
	app.Get("/tasks", controllers.GetAllTasks)
	app.Get("/tasks/:id", controllers.GetTask)
	app.Post("/tasks", controllers.CreateTask)
	app.Delete("/tasks/:id", controllers.DeleteTask)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi World! Welcome!")
	})

	app.Listen(":3000")
}
