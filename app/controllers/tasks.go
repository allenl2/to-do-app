package controllers

import (
	"log"
	"strconv"

	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/gofiber/fiber/v2"
)

//returns all tasks
func GetAllTasks(c *fiber.Ctx) error {
	var tasks []models.Task

	//search for all tasks
	result := database.RetrieveAllTasks(&tasks)

	if result.Error != nil {
		log.Println("Unable to retrieve tasks.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    tasks,
	})
}

//returns the tasks with the specified id
func GetTask(c *fiber.Ctx) error {
	var task models.Task
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		log.Println("Invalid task.", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task id.")
	}

	//search for the tasks
	result := database.RetrieveTask(&task, uint(id))

	if result.Error != nil {
		log.Println("Unable to find task.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    task,
	})
}

//creates a new task with info provided in the body
func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	err := c.BodyParser(task)

	if err != nil {
		log.Println("Unable to create new task from data.")
		return c.Status(fiber.StatusBadRequest).SendString("Unable to create new task from data.")
	}

	result := database.CreateNewTask(task)

	if result.Error != nil {
		log.Println("Unable to create new task in database.", result.Error.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    task,
	})
}

//deletes the task with the specified id
func DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	task := &models.Task{ID: uint(id)}

	if err != nil {
		log.Println("Invalid task id.")
		return c.Status(fiber.StatusNotFound).SendString("Invalid task id.")
	}

	result := database.DeleteTask(task)

	if result.Error != nil {
		log.Println("Unable to delete task.", result.Error.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.Status(fiber.StatusOK).SendString("Task deleted.")
}
