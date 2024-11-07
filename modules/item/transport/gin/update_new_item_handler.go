package ginitem

import (
	"net/http"
	"strconv"
	"todo/common"
	"todo/modules/item/biz"
	item "todo/modules/item/model"
	"todo/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData item.TodoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindJSON(&itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSQLStorage(db)
		biz := biz.NewUpdateItemBiz(store)

		if err := biz.UpdateItem(c.Request.Context(), map[string]interface{}{"id": id}, &itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData))
	}
}
