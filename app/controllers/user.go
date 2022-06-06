package controllers

import (
	"log"
	"strconv"
	"to-do-app/app/database"
	"to-do-app/app/models"
	"to-do-app/app/utils"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

//searches for the user based on id
func GetUser(c *fiber.Ctx) error {
	var user models.User
	var resUser models.UserResponse
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		log.Println("Invalid user id.", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid user id.",
		})
	}

	//search for the user with the given id
	result := database.RetrieveUser(&user, uint(id))

	if result.Error != nil {
		log.Println("User not found.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": result.Error.Error(),
		})
	}

	if err := copier.Copy(&resUser, &user); err != nil {
		log.Println("Unable to return user. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
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

	if parseErr := c.BodyParser(&user); parseErr != nil {
		log.Println("Unable to create new user.", parseErr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": parseErr.Error(),
		})
	}

	if err := utils.ValidateUserPassword(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input, cannot parse data.",
		})
	}

	hash, hashErr := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if hashErr != nil {
		log.Println("Unable to create new user.", hashErr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": hashErr.Error(),
		})
	}

	user.Password = hash

	//add the user to the database
	result := database.CreateNewUser(&user)

	if result.Error != nil {
		log.Println("Unable to add user.", result.Error.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": result.Error.Error(),
		})
	}

	if err := copier.Copy(&resUser, &user); err != nil {
		log.Println("Unable to add user. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    resUser,
	})
}

//updates the details of the specified user
func UpdateUser(c *fiber.Ctx) error {
	var dbUser models.User
	var inputUser models.User
	var resUser models.UserResponse

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	parseErr := c.BodyParser(&inputUser)

	if err != nil || parseErr != nil {
		log.Println("Invalid input.", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input, cannot parse data.",
		})
	}

	if err := utils.ValidateStruct(inputUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input, cannot parse data.",
		})
	}

	//search for the user with the given id
	if res := database.RetrieveUser(&dbUser, uint(id)); res.Error != nil {
		log.Println("User not found.", res.Error.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": res.Error.Error(),
		})
	}

	//update any fields that are provided
	if inputUser.Username != "" {
		dbUser.Username = inputUser.Username
	}
	if inputUser.Password != "" {
		hash, hashErr := argon2id.CreateHash(inputUser.Password, argon2id.DefaultParams)
		if hashErr != nil {
			log.Println("Unable to update user.", hashErr)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Unable to update user.",
			})
		}
		dbUser.Password = hash
	}

	//save the changes to the DB
	if res := database.UpdateUser(&dbUser); res.Error != nil {
		log.Println("Unable to update user.", res.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": res.Error.Error(),
		})
	}

	//create response object
	if err := copier.Copy(&resUser, &dbUser); err != nil {
		log.Println("Unable to update user. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    resUser,
	})
}

//renders the account details on the Account page
func RenderAccount(c *fiber.Ctx) error {
	var user models.User
	currSess, sessErr := database.SessionStore.Get(c)

	if sessErr != nil {
		log.Println("Unable to get user info.")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Unable to get user info.",
		})
	}

	result := database.RetrieveUser(&user, currSess.Get("userID").(uint))

	if result.Error != nil {
		log.Println("Unable to retrieve user.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": result.Error.Error(),
		})
	}

	return c.Render("account", fiber.Map{
		"Username": user.Username,
		"UserID":   user.ID,
	})
}
