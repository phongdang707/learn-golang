package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func main() {
	now := time.Now()
	item := Item{
		Title:       "Test Item",
		Description: "This is a test item",
		Status:      "active",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	json, err := json.Marshal(item)
	if err != nil {
		log.Fatalf("Error marshalling item: %v", err)
	}

	r := gin.Default()
	r.Group("v1")
	{
		r.Group("items")
		{
			r.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{"data": string(json)})
			})
			r.POST("/", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Hello, World!"})
			})
			r.GET("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Hello, World!"})
			})
			r.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Hello, World!"})
			})
			r.DELETE("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Hello, World!"})
			})
		}
	}
	r.Run(":8080")
}
