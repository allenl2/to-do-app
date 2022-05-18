package controllers

import (
	"log"
	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

// func LoginAuth(username, pass string) bool {
// 	//To-DO: store something in the session
// 	//To-Do: allow for logout
// 	var user models.User

// 	//search for the user in DB based on username
// 	result := database.RetrieveUser(&user, username)

// 	if result.Error != nil {
// 		log.Println("User not found.", result.Error.Error()) //To-do: standardize login error messages.
// 		return false
// 	}

// 	//check if the entered password matches the hashed password
// 	match, err := argon2id.ComparePasswordAndHash(pass, user.Password)
// 	if err != nil {
// 		log.Println("Error. Unable to verify login info.", err)
// 		return false
// 	}

// 	if match {
// 		log.Println("Successful login.")
// 		return true
// 	}

// 	log.Println("Invalid credentials.")
// 	return false
// }

func LoginAuth(c *fiber.Ctx) error {
	var user models.User
	parseErr := c.BodyParser(&user)
	enteredPass := user.Password

	//search for the user in DB based on username
	result := database.RetrieveUser(&user, user.Username)

	//check if the entered password matches the hashed password
	match, hashErr := argon2id.ComparePasswordAndHash(enteredPass, user.Password)

	if parseErr != nil {
		log.Println("Error. Invalid credentials.", parseErr)
		return fiber.NewError(fiber.StatusUnauthorized, parseErr.Error())
	}
	if hashErr != nil {
		log.Println("Error. Invalid credentials.", hashErr)
		return fiber.NewError(fiber.StatusUnauthorized, hashErr.Error())
	}
	if result.Error != nil {
		log.Println("Error. Invalid credentials.", result.Error)
		return fiber.NewError(fiber.StatusUnauthorized, result.Error.Error())
	}
	//if successful login, create new session and send cookie
	if match {
		log.Println("Successful login.")

		//creates new session
		currSess, err := database.SessionStore.Get(c)
		if err != nil {
			log.Println("Error creating new session", err)
			return err
		}

		//creates new session_id
		if err := currSess.Regenerate(); err != nil {
			log.Println("Error regenerating new session id", err)
			return err
		}

		currSess.Set("user", user.Username)

		log.Println("current session id: ", currSess.ID())
		log.Println("current session keys: ", currSess.Keys())
		log.Println("current session name: ", currSess.Get("user"))

		currSess.Save()

		return nil
	}

	return fiber.NewError(fiber.StatusUnauthorized)
}

func CheckAuth(c *fiber.Ctx) error {

	//get the session from the current user request
	currSess, err := database.SessionStore.Get(c)
	if err != nil {
		log.Println("Error. Please login again.", err)
		return err
	}

	//checks if this session is authenticated and belongs to a user, if so continue
	if currSess.Get("user") != nil {
		log.Println("Session is valid and authenticated")
		return c.Next()
	}

	//if not logged in, redirect to login page
	log.Println("Not logged in. Please login first")
	return c.Redirect("/")
}

//to-do
//add logout feature
