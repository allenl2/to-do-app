package database

import (
	"log"
	"to-do-app/app/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() *gorm.DB {
	//get environment variables for DB
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Unable to read environment variables")
	}

	dbUser := viper.Get("dbUser").(string)
	dbPass := viper.Get("dbPass").(string)

	dbURL := "host=localhost dbname=todo port=5432 " + "user=" + dbUser + " password=" + dbPass
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
func AutoMigrateDB() {

	autoMigErr := DB.AutoMigrate(&models.Task{}, &models.User{})
	if autoMigErr != nil {
		log.Println("Error occurred while auto migrating database.")
	}
}

//User DB Functions

func RetrieveUser(user *models.User, id uint) *gorm.DB {
	return DB.Where(&models.User{ID: id}).Preload("Tasks").First(user)
}

func RetrieveUserByUsername(user *models.User, username string) *gorm.DB {
	return DB.Where(&models.User{Username: username}).First(user)
}

func CreateNewUser(user *models.User) *gorm.DB {
	return DB.Create(user)
}

func UpdateUser(user *models.User) *gorm.DB {
	return DB.Save(user)
}

//Tasks DB Functions

func RetrieveAllTasks(tasks *[]models.Task) *gorm.DB {
	return DB.Find(tasks)
}

func RetrieveTask(task *models.Task, id uint) *gorm.DB {
	return DB.Where(&models.Task{ID: id}).First(task)
}

func CreateNewTask(task *models.Task) *gorm.DB {
	return DB.Create(task)
}

func DeleteTask(id uint) *gorm.DB {
	return DB.Delete(&models.Task{}, id)
}

func UpdateTask(task *models.Task) *gorm.DB {
	return DB.Save(task)
}
