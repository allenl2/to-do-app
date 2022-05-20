package controllers

import (
	"log"
	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

//searches for the user based on username
func GetUser(c *fiber.Ctx) error {
	var user models.User
	var resUser models.UserResponse

	//search for the user with the given username
	result := database.RetrieveUser(&user, c.Params("username"))

	if result.Error != nil {
		log.Println("User not found.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}

	if err := copier.Copy(&resUser, &user); err != nil {
		log.Println("Unable to return user. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    resUser,
	})
}

//creates a new user with the provided username, default password
func CreateUser(c *fiber.Ctx) error {

	var user models.User
	var resUser models.UserResponse
	parseErr := c.BodyParser(&user)

	if parseErr != nil {
		log.Println("Unable to create new user.", parseErr)
		return c.Status(fiber.StatusBadRequest).SendString(parseErr.Error())
	}

	hash, hashErr := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if hashErr != nil {
		log.Println("Unable to create new user.", hashErr)
		return c.Status(fiber.StatusInternalServerError).SendString(hashErr.Error())
	}

	user.Password = hash

	//add the user to the database
	result := database.CreateNewUser(&user)

	if result.Error != nil {
		log.Println("Unable to add user.", result.Error.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	if err := copier.Copy(&resUser, &user); err != nil {
		log.Println("Unable to add user. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    resUser,
	})
}
