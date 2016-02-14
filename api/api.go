package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db gorm.DB

func main() {
	var err error

	db, err = gorm.Open("sqlite3", "./blog.db")

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Post{})

	router := gin.Default()

	router.POST("/users", handleNewUsers)
	router.POST("/posts", handleNewPosts)
	router.GET("/posts", handlePosts)
	router.GET("/posts/:id", handleSinglePost)
	router.PATCH("/posts/:id", handleUpdatePost)

	log.Fatal(router.Run(":8080"))
}
