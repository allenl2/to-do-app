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

func GetUser(c *fiber.Ctx) error {
	var user models.User
	log.Println("Made it to the controller")

	//search for the user with the given username (Hard coded)
	result := database.FindUser(&user, "John")

	if result.Error != nil {
		log.Println("User not found.", result)
		return result.Error
	}
	return c.JSON(models.User{
		Username: user.Username,
		Password: user.Password,
	})
}
