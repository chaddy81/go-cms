package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/blackfriday"
	"log"
	"strconv"
	"time"
)

type Model struct {
	Id        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string `sql:"unique"`
	Posts     []Post // One-To-Many relationship (has many)
}

type Post struct {
	gorm.Model
	Title    string `sql:"unique"`
	HTML     string
	Markdown string
	UserID   int `sql:"index"`
}

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

	router.POST("/users/new", handleNewUsers)
	router.POST("/posts/new", handleNewPosts)
	router.GET("/posts", handlePosts)
	router.GET("/posts/:id", handleSinglePost)
	router.PATCH("/posts/:id", handleUpdatePost)

	log.Fatal(router.Run(":8080"))
}

func handleNewUsers(c *gin.Context) {

	user := User{
		Email:     c.Query("email"),
		Firstname: c.Query("firstname"),
		Lastname:  c.Query("lastname"),
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(422, gin.H{"status": 422, "message": err.Error()})
		return
	}

	c.JSON(201, user)
}

func handleNewPosts(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("userid"))
	html := blackfriday.MarkdownBasic([]byte(c.PostForm("markdown")))

	post := Post{
		Title:    c.PostForm("title"),
		HTML:     string(html),
		Markdown: c.PostForm("markdown"),
		UserID:   userId,
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(422, gin.H{"status": 422, "message": err.Error()})
		return
	}

	c.JSON(201, post)
}

func handleUpdatePost(c *gin.Context) {
	var post Post
	id := c.Param("id")

	db.First(&post, id)

	html := blackfriday.MarkdownBasic([]byte(c.PostForm("markdown")))

	post.Title = c.PostForm("title")
	post.HTML = string(html)
	post.Markdown = c.PostForm("markdown")

	if err := db.Save(&post).Error; err != nil {
		c.JSON(422, gin.H{"status": 422, "message": err.Error()})
		return
	}

	c.JSON(201, post)
}

func handlePosts(c *gin.Context) {
	posts := []Post{}

	if err := db.Find(&posts).Error; err != nil {
		c.JSON(422, gin.H{"status": 422, "message": err.Error()})
		return
	}

	c.JSON(200, posts)
}

func handleSinglePost(c *gin.Context) {
	var post Post
	id := c.Param("id")

	if err := db.First(&post, id).Error; err != nil {
		c.JSON(422, gin.H{"status": 422, "message": err.Error()})
		return
	}

	c.JSON(200, post)
}
