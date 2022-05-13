package controllers

import (
	"log"
	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/gofiber/fiber/v2"
)

// func CreateUser(c *fiber.Ctx) error {
// 	user := models.User{Username: "John Smith", Password: "password1"}
// 	var db *gorm.DB
// 	result := db.Create(&user)

// 	if result != nil {
// 		log.Println("Failed to create new user.")
// 		return nil
// 	}

// 	log.Println("Created new user.")
// 	return nil
// }

// func GetUsers(c *fiber.Ctx) error {
// 	var user models.User

// }

//searches for the user based on username
func GetUser(c *fiber.Ctx) error {
	var user models.User

	//search for the user with the given username
	result := database.RetrieveUser(&user, c.Params("username"))

	if result.Error != nil {
		log.Println("User not found.", result)
		return result.Error
	}
	return c.JSON(models.User{
		Username: user.Username,
		Password: user.Password,
	})
}

//creates a new user with the provided username, default password
func CreateUser(c *fiber.Ctx) error {
	user := models.User{Username: c.Params("username"), Password: "password"}

	//search for the user with the given username
	result := database.CreateNewUser(&user)

	if result.Error != nil {
		log.Println("Unable to add user.", result)
		return result.Error
	}
	return c.JSON(models.User{
		Username: user.Username,
		Password: user.Password,
	})
}
