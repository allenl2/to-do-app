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

	//Initialize the database
	database.Init()
	autoMigErr := database.AutoMigrateDB()

	if autoMigErr != nil {
		log.Println("Error occurred while auto migrating database")
	}

	//Initialize redis cache and session
	database.InitSession()

	//Authenication middelware
	// app.Use("/tasks", basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"admin": "123456",
	// 	},
	// 	Realm:      "Forbidden",
	// 	Authorizer: controllers.LoginAuth,
	// }))

	//PUBLIC API ROUTES
	apiPublic := app.Group("/api")
	apiPublic.Post("/login", controllers.LoginAuth)

	//PRIVATE API ROUTES
	api := app.Group("/api", controllers.CheckAuth)

	//Users
	api.Get("/user/:username", controllers.GetUser)
	api.Post("/api/user", controllers.CreateUser)

	//Tasks
	api.Get("/tasks", controllers.GetAllTasks)
	api.Get("/tasks/:id", controllers.CheckAuth, controllers.GetTask)
	api.Post("/tasks", controllers.CreateTask)
	api.Delete("/tasks/:id", controllers.DeleteTask)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi World! Welcome!")
	})

	app.Listen(":3000")
}
