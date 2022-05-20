package controllers

import (
	"log"
	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func LoginAuth(c *fiber.Ctx) error {
	var user models.User
	parseErr := c.BodyParser(&user)
	enteredPass := user.Password

	//search for the user in DB based on username
	result := database.RetrieveUserByUsername(&user, user.Username)

	//check if the entered password matches the hashed password
	match, hashErr := argon2id.ComparePasswordAndHash(enteredPass, user.Password)

	if parseErr != nil || hashErr != nil || result.Error != nil {
		log.Println("Error. Invalid credentials.")
		return c.Status(fiber.StatusUnauthorized).SendString("Error. Invalid credentials.")
	}
	//if successful login, create new session and send cookie
	if match {
		log.Println("Successful login.")

		//creates new session
		currSess, err := database.SessionStore.Get(c)
		if err != nil {
			log.Println("Error creating new session", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		//creates new session_id
		if err := currSess.Regenerate(); err != nil {
			log.Println("Error regenerating new session id", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		currSess.Set("user", user.Username)
		sessErr := currSess.Save()
		if sessErr != nil {
			log.Println("Error saving session", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).SendString("Success.")
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Error. Invalid credentials.")
}

func CheckAuth(c *fiber.Ctx) error {

	//get the session from the current user request
	currSess, err := database.SessionStore.Get(c)
	if err != nil {
		log.Println("Error. Please login again.", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error. Please try again.")
	}

	//checks if this session is authenticated and belongs to a user, if so continue
	if currSess.Get("user") != nil {
		return c.Next()
	}

	//if not logged in, redirect to login page
	log.Println("Not logged in. Please login first.")
	return c.Status(fiber.StatusUnauthorized).Redirect("/")
}
