package main

import (
	"fmt"
	"log"
	"os"
	"time"
	ginitem "todo/modules/item/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dsn := os.Getenv("DB_CONNECTION")
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	db = db.Debug()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("Database connected", db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			// items.GET("/", ListItem(db))
			items.POST("/", ginitem.CreateItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			// items.PUT("/:id", UpdateItem(db))
			// items.DELETE("/:id", DeleteItem(db))
		}
	}
	r.Run(":8080")
}



// func GetItem(db *gorm.DB) func(ctx *gin.Context) {
// 	return func(c *gin.Context) {
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var item Item
// 		if err := db.Where("id = ?", id).First(&item).Error; err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse(item))
// 	}
// }

// func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
// 	return func(c *gin.Context) {
// 		var itemData TodoItemUpdate
		
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if err := c.ShouldBindJSON(&itemData); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if err := db.Where("id = ?", id).Updates(&itemData).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
// 	}
// }

// func DeleteItem(db *gorm.DB) func(ctx *gin.Context) {
// 	return func(c *gin.Context) {
// 		now := time.Now()
// 		var itemData TodoItemDelete
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		status := "Deleted"
// 		itemData.Id = id
// 		itemData.DeletedAt = &now
// 		itemData.Status = &status
// 		if err := db.Where("id = ?", id).Updates(&itemData).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
// 	}
// }

// func ListItem (db *gorm.DB) func(ctx *gin.Context) {
// 	return func(c *gin.Context) {
// 		var paging common.Paging
// 		var result []Item
// 		if err := c.ShouldBindQuery(&paging); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		paging.Process()	
// 		if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Count(&paging.Total).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
// 	}
// }
