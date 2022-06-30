package database

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/Antuan01/go-test/models"
	"gorm.io/driver/sqlite"
)

var Database *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
	  panic("failed to connect database")
	}
	
	fmt.Println("Connected to the database!")

	db.AutoMigrate(&models.Post{}, &models.Comment{}, &models.PostReport{}, &models.CommentReport{})

	db.Create(&models.Post{Contents: "Post 2"})
	db.Create(&models.Post{Contents: "Post 3"})
	db.Create(&models.Post{Contents: "Post 1"})

	db.Create(&models.Comment{Contents: "Comment 2"})
	db.Create(&models.Comment{Contents: "Comment 3"})
	db.Create(&models.Comment{Contents: "Comment 1"})

	fmt.Println("Database Migrated!")

	Database = db
}