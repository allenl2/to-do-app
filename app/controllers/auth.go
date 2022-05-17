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
		sess := database.CreateSession()

		//creates a cookie with session_id
		sessCookie := new(fiber.Cookie)
		sessCookie.Name = sess.KeyLookup
		sessCookie.Value = sess.KeyGenerator()

		c.Cookie(sessCookie)

		return nil
	}

	return fiber.NewError(fiber.StatusUnauthorized)
}
