package controllers

import (
	"log"
	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/alexedwards/argon2id"
)

func LoginAuth(username, pass string) bool {
	//To-DO: store something in the session
	//To-Do: allow for logout
	var user models.User

	//search for the user in DB based on username
	result := database.RetrieveUser(&user, username)

	if result.Error != nil {
		log.Println("User not found.", result.Error.Error()) //To-do: standardize login error messages.
		return false
	}

	//check if the entered password matches the hashed password
	match, err := argon2id.ComparePasswordAndHash(pass, user.Password)
	if err != nil {
		log.Println("Error. Unable to verify login info.", err)
		return false
	}

	if match {
		log.Println("Successful login.")
		return true
	}

	log.Println("Invalid credentials.")
	return false
}
