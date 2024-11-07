package ginitem

import (
	"net/http"
	"todo/common"
	"todo/modules/item/biz"
	item "todo/modules/item/model"
	"todo/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData item.TodoItemCreation
		if err := c.ShouldBindJSON(&itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStorage(db)
		biz := biz.NewCreateNewItemBiz(store)

		if err := biz.CreateNewItem(c.Request.Context(), &itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData.Id))
	}
}