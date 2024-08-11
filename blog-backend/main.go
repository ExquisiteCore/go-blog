package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Post{})

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/posts", func(c *gin.Context) {
		var post Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&post)
		c.JSON(http.StatusOK, post)
	})

	r.GET("/posts/:id", func(c *gin.Context) {
		var post Post
		if err := db.First(&post, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		htmlContent := blackfriday.Run([]byte(post.Content))
		c.JSON(http.StatusOK, gin.H{"title": post.Title, "content": string(htmlContent)})
	})

	r.Run(":8080")
}
