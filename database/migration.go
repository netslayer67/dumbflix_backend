package database

import (
	"fmt"
	"dumbflix/models"
	"dumbflix/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Profile{},
		&models.User{},
		&models.Category{},
		&models.Film{},
		&models.Episode{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
