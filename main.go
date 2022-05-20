package main

import (
	"log"
	"to-do-app/app/controllers"
	"to-do-app/app/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Setup the server using Go Fiber
	app := fiber.New()

	//Initialize the database & session
	database.Init()
	database.AutoMigrateDB()
	database.InitSession()

	//PUBLIC API ROUTES
	apiPublic := app.Group("/api")
	apiPublic.Post("/login", controllers.LoginAuth)
	apiPublic.Post("/user", controllers.CreateUser)

	//PRIVATE API ROUTES
	api := app.Group("/api", controllers.CheckAuth)

	//Users
	api.Get("/user/:username", controllers.GetUser)

	//Tasks
	api.Get("/tasks", controllers.GetAllTasks)
	api.Get("/tasks/:id", controllers.GetTask)
	api.Post("/tasks", controllers.CreateTask)
	api.Delete("/tasks/:id", controllers.DeleteTask)

	//Base
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi World! Welcome to the Home page!")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Println("Server failed to launch.")
	}
}
