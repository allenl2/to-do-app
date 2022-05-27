package controllers

import (
	"log"
	"strconv"

	"to-do-app/app/database"
	"to-do-app/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

//returns all tasks
func GetAllTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	var resTasks []models.TaskResponse

	//search for all tasks
	result := database.RetrieveAllTasks(&tasks)

	if result.Error != nil {
		log.Println("Unable to retrieve tasks.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}

	if err := copier.Copy(&resTasks, &tasks); err != nil {
		log.Println("Unable to retrieve tasks. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    resTasks,
	})
}

//returns the tasks with the specified id
func GetTask(c *fiber.Ctx) error {
	var task models.Task
	var resTask models.TaskResponse
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		log.Println("Invalid task.", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task id.")
	}

	//search for the task
	result := database.RetrieveTask(&task, uint(id))

	if result.Error != nil {
		log.Println("Unable to find task.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}

	if err := copier.Copy(&resTask, &task); err != nil {
		log.Println("Unable to retrieve task. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    resTask,
	})
}

//creates a new task with info provided in the body, associated with current user
func CreateTask(c *fiber.Ctx) error {

	task := new(models.Task)
	err := c.BodyParser(task)
	var resTask models.TaskResponse

	if err != nil {
		log.Println("Unable to create new task from data.")
		return c.Status(fiber.StatusBadRequest).SendString("Unable to create new task from data.")
	}

	currSess, sessErr := database.SessionStore.Get(c)

	if sessErr != nil {
		log.Println("Unable to create new task from data. User session error.")
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to create new task from data.  User session error.")
	}

	task.UserID = currSess.Get("userID").(uint)
	task.Assignee = currSess.Get("username").(string)

	result := database.CreateNewTask(task)

	if result.Error != nil {
		log.Println("Unable to create new task in database.", result.Error.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	if err := copier.Copy(&resTask, &task); err != nil {
		log.Println("Unable to create task. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    resTask,
	})
}

//deletes the task with the specified id
func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		log.Println("Invalid task id.")
		return c.Status(fiber.StatusNotFound).SendString("Invalid task id.")
	}

	result := database.DeleteTask(uint(id))

	if result.Error != nil {
		log.Println("Unable to delete task.", result.Error.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		log.Println("Task not found.")
		return c.Status(fiber.StatusNotFound).SendString("Task not found.")
	}

	return c.Status(fiber.StatusOK).SendString("Task deleted.")
}

//update the details of the specified task
func UpdateTask(c *fiber.Ctx) error {
	var dbTask models.Task
	var inputTask models.Task
	var resTask models.TaskResponse

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	parseErr := c.BodyParser(&inputTask)

	if err != nil || parseErr != nil {
		log.Println("Invalid input.", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input.")
	}

	//search for the task with the given id
	if res := database.RetrieveTask(&dbTask, uint(id)); res.Error != nil {
		log.Println("Task not found.", res.Error.Error())
		return c.Status(fiber.StatusNotFound).SendString(res.Error.Error())
	}

	//update any fields that are provided
	if inputTask.TaskName != "" {
		dbTask.TaskName = inputTask.TaskName
	}
	if inputTask.Assignee != "" {
		dbTask.Assignee = inputTask.Assignee
	}
	if inputTask.Status != "" {
		dbTask.Status = inputTask.Status
	}

	//save the changes to the DB
	if res := database.UpdateTask(&dbTask); res.Error != nil {
		log.Println("Unable to update task.", res.Error.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	//create response object
	if err := copier.Copy(&resTask, &dbTask); err != nil {
		log.Println("Unable to update task. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    resTask,
	})
}

func RenderTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	var resTasks []models.TaskResponse

	result := database.RetrieveAllTasks(&tasks)
	if result.Error != nil {
		log.Println("Unable to retrieve tasks.", result.Error.Error())
		return c.Status(fiber.StatusNotFound).SendString(result.Error.Error())
	}

	if err := copier.Copy(&resTasks, &tasks); err != nil {
		log.Println("Unable to retrieve tasks. Copying error.", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render("home", fiber.Map{
		"Tasks": resTasks,
	})
}
