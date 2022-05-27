package main

import (
	"log"
	"to-do-app/app/controllers"
	"to-do-app/app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	//initialize Go html template engine
	engine := html.New("./app/views", ".html")

	//Setup the server using Go Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Initialize the database & session
	database.Init()
	database.AutoMigrateDB()
	database.InitSession()

	//PUBLIC API ROUTES
	apiPublic := app.Group("/api")
	apiPublic.Post("/login", controllers.LoginAuth)
	apiPublic.Post("/logout", controllers.LogoutAuth)
	apiPublic.Post("/user", controllers.CreateUser)

	//PRIVATE API ROUTES
	api := app.Group("/api", controllers.CheckAuth)

	//Users
	api.Get("/user/:id", controllers.GetUser)
	api.Patch("/user/:id", controllers.UpdateUser)

	//Tasks
	api.Get("/tasks", controllers.GetAllTasks)
	api.Get("/tasks/:id", controllers.GetTask)
	api.Post("/tasks", controllers.CreateTask)
	api.Delete("/tasks/:id", controllers.DeleteTask)
	api.Patch("/tasks/:id", controllers.UpdateTask)

	//Static files
	app.Static("/static", "./static")

	//Base
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/home", controllers.CheckAuth, controllers.RenderTasks)
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{})
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Println("Server failed to launch.")
	}
}
