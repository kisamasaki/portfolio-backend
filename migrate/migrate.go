package main

import (
	"fmt"
	"portfolio-backend/db"
	"portfolio-backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer func() {
		if sqlDB, err := dbConn.DB(); err == nil {
			if err := sqlDB.Close(); err != nil {
				fmt.Printf("Error closing DB: %v", err)
			}
		}
	}()

	if err := dbConn.AutoMigrate(&model.User{}, &model.UserAuth{}, &model.Comic{}, &model.Following{}, &model.CompletedComic{}); err != nil {
		fmt.Printf("Error migrating database: %v", err)
		return
	}

	fmt.Println("Successfully Migrated")
}
