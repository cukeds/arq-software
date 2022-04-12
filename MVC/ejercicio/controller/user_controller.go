package controller

import (
	"git/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context){

	var user1 dto.User
	var user2 dto.User

	user1.Name = "Jose"
	user2.Name = "Pancho"
	users := [2]dto.User{user1, user2}
	c.JSON(http.StatusOK, users)
}
