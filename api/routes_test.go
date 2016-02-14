package main

import (
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	// "net/http"
	// "net/http/httptest"
)

var _ = Describe("Post", func() {

	var (
		body []byte
		err  error
	)

	BeforeEach(func() {
		var err error

		db, err = gorm.Open("sqlite3", "./blog_test.db")
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&User{}, &Post{})

		// rec = httptest.NewRecorder()
	})

	AfterEach(func() {
		db.DropTable(&User{})
		db.DropTable(&Post{})
		db.Close()
	})

	Context("List all posts", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/posts", handlePosts)
			Expect(response.Code).To(Equal(200))
			log.Print(response)
		})
	})

	Context("Create a Post", func() {

		BeforeEach(func() {
			user := User{
				Firstname: "Chad",
				Lastname:  "Bartels",
				Email:     "chad.bartels@gmail.com",
			}

			if err := db.Create(&user).Error; err != nil {
				log.Println("Error creating User")
			}

			post := Post{
				Title:    "Test Title",
				Markdown: "h1 test\r\n=======",
				UserID:   3,
			}

			body, err = json.Marshal(post)
			if err != nil {
				log.Println("Unable to marshal post")
			}
		})

		It("returns a 201 Status Code", func() {
			PostRequest("POST", "/posts/new", handleNewPosts, bytes.NewReader(body))
			Expect(response.Code).To(Equal(201))
		})
	})
})
