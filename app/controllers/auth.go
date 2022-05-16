package controllers

import (
	"log"
	"to-do-app/app/database"
	"to-do-app/app/models"
)

func LoginAuth(username, pass string) bool {
	//implement hashing here?
	var user models.User

	//search for the user in DB based on username
	result := database.RetrieveUser(&user, username)

	if result.Error != nil {
		log.Println("User not found.", result.Error.Error())
		return false
	}

	//if the password matches
	if pass == user.Password {
		return true
	}

	return false
}
