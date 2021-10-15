package repository

import (
	"main/helper"
	"main/service"

	"github.com/gin-gonic/gin"
)

func GetUserBookingBank(c *gin.Context) {
	userId := c.Param("id")
	resUsers, err := service.GetUserBookingBank(c, userId)
	helper.PanicIfError(err)
	c.JSON(200, gin.H{
		"data": resUsers["data"],
	})
}