package routes

import (
	"net/http"
	"simple-api-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi..!",
		})
	})

	routes_api := r.Group("/api/v1")
	{
		user := new(controllers.User)
		routes_api.GET("/users", user.Index)
		routes_api.POST("/users", user.Store)
		routes_api.PUT("/users", user.Update)
		routes_api.DELETE("/users/:id", user.Destroy)
		routes_api.GET("/users/:id", user.Show)
	}

	return r
}

func Run() *gin.Engine {
	r := SetupRouter()
	r.Run(":8081")

	return r
}
