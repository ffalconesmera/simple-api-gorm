package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Response struct {
}

func (response *Response) Json(c *gin.Context, result *gorm.DB, data interface{}) {
	if result.Error == nil {
		response.JsonSuccess(c, http.StatusOK, gin.H{"data": data})
		return
	}
	response.JsonFail(c, http.StatusBadRequest, result.Error.Error())
	return
}

func (response *Response) JsonSuccess(c *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	c.JSON(status, h)
	return
}

func (response *Response) JsonFail(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  "fail",
		"message": message,
		"data":    nil,
		"code":    status,
	})
}
