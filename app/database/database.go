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
	} else {
		log.Println("Connected to database!")
	}

	return DB
}

//auto migrates the models into the database
func AutoMigrateDB() error {
	//check that db connection is working
	if err != nil {
		return err
	}

	autoMigErr := DB.AutoMigrate(&models.Task{}, &models.User{})
	return autoMigErr
}

func FindUser(user *models.User, username string) *gorm.DB {
	log.Println(DB.First(user).Error)
	//return DB.First(user)
	return DB.Where(&models.User{Username: "John"}).First(user)
}
