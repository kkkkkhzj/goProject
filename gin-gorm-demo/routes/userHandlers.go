package routes

import (
	"gorm-gin-test/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func createUser(c *gin.Context) {
	var req createAUserRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind" + err.Error(),
		})
	}

	if err := model.GetCurrentDB().InsertAUser(req.Username, req.Password, req.Email); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot insert user into database" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

type userInfo struct {
	UserID   uint   `json:"user_id" `
	Username string `json:"username" `
	Email    string `json:"email"`
}

func listUsers(c *gin.Context) {
	var usersResponse []userInfo
	users, err := model.GetCurrentDB().QueryAllUsers()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get users from database" + err.Error(),
		})
	}

	for _, u := range users {
		usersResponse = append(usersResponse, userInfo{
			UserID:   u.ID,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": usersResponse,
	})
}

// NOTE: 作为练习
func listUserByID(c *gin.Context) {
	// 可参考 https://github.com/gin-gonic/gin#bind-uri
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
