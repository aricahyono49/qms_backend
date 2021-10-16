package repository

import (
	"main/helper"
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user  model.User
	errA := c.ShouldBind(&user)
	helper.PanicIfError(errA)
	resUsers, err := service.Login(c, user)
	helper.PanicIfError(err)
	status := resUsers["status"].(int)
	c.JSON(status, gin.H{
		"message": resUsers["message"],
		"user_id": resUsers["user_id"],

	})
}

func Register(c *gin.Context) {
	var user  model.User
	err := c.ShouldBind(&user)
	helper.PanicIfError(err)

	resUsers, err := service.Register(c, user)
	helper.PanicIfError(err)

	status := resUsers["status"].(int)
	c.JSON(status, gin.H{
		"message": resUsers["message"],
	})
}