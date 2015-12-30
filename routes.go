package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/mattn/go-sqlite3"
  "github.com/russross/blackfriday"
  // "log"
  "strconv"
)

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

  // c.Request.ParseMultipartForm(32 << 20)
  // file, handler, err := c.Request.FormFile("header")

  // if err != nil {
  //   log.Print(err)
  // }

  // defer file.Close()

  // log.Print(handler.Header)

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
