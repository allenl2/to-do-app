package controllers

import (
	"log"
	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

//searches for the user based on username
func GetUser(c *fiber.Ctx) error {
	var user models.User

	//search for the user with the given username
	result := database.RetrieveUser(&user, c.Params("username"))

	if result.Error != nil {
		log.Println("User not found.", result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.JSON(models.User{
		Username: user.Username,
		Password: user.Password,
	})
}

//creates a new user with the provided username, default password
func CreateUser(c *fiber.Ctx) error {

	var user models.User
	parseErr := c.BodyParser(&user)

	if parseErr != nil {
		log.Println("Unable to create new user.", parseErr)
		return fiber.NewError(fiber.StatusBadRequest, parseErr.Error())
	}

	hash, hashErr := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if hashErr != nil {
		log.Println("Unable to create new user.", hashErr)
		return fiber.NewError(fiber.StatusInternalServerError, hashErr.Error())
	}

	user.Password = hash

	//add the user to the database
	result := database.CreateNewUser(&user)

	if result.Error != nil {
		log.Println("Unable to add user.", result.Error.Error())
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(models.User{
		Username: user.Username,
		Password: user.Password,
	})
}
