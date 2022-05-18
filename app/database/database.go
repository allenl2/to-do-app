package database

import (
	"log"
	"to-do-app/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() *gorm.DB {
	dbURL := "host=localhost user=todo password=secret dbname=todo port=5432"
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Println(err)
		return nil
	} else {
		log.Println("Connected to database!")
		return DB
	}
}

//auto migrates the models into the database
func AutoMigrateDB() error {

	autoMigErr := DB.AutoMigrate(&models.Task{}, &models.User{})
	if autoMigErr != nil {
		log.Println("Error occurred while auto migrating database.")
		return autoMigErr
	}
	return nil
}

func RetrieveUser(user *models.User, username string) *gorm.DB {
	return DB.Where(&models.User{Username: username}).First(user)
}

func CreateNewUser(user *models.User) *gorm.DB {
	return DB.Create(user)
}

func RetrieveAllTasks(tasks *[]models.Task) *gorm.DB {
	return DB.Find(tasks)
}

func RetrieveTask(task *models.Task, id uint) *gorm.DB {
	return DB.Where(&models.Task{ID: id}).Find(task)
}

func CreateNewTask(task *models.Task) *gorm.DB {
	return DB.Create(task)
}

func DeleteTask(task *models.Task) *gorm.DB {
	return DB.Delete(task)
}
