package database

import (
	"housy/models"
	"housy/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.ListAs{},
		// &models.Profile{},
		// &models.Product{},
		// &models.Category{},
		// &models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}