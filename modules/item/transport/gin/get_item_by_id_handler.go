package ginitem

import (
	"net/http"
	"strconv"
	"todo/common"
	"todo/modules/item/biz"
	"todo/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := storage.NewSQLStorage(db)
		biz := biz.NewGetItemByIdBiz(storage)

		item, err := biz.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(item))
	}
}
