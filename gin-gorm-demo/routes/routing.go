package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", createUser)
	r.GET("/users", listUsers)
	r.GET("/users/:userID", listUserByID)
	return r
}
