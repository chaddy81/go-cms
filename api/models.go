package main

import (
	"github.com/jinzhu/gorm"
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
	Title    string `sql:"unique" json:"title`
	HTML     string `json:"html"`
	Markdown string `json:"markdown"`
	UserID   int    `sql:"index" json:"userid`
}
