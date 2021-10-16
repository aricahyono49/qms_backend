package repository

import (
	"fmt"
	"main/helper"
	"main/model"
	"main/service"
	"time"

	"github.com/gin-gonic/gin"
)

func BankIndex(c *gin.Context) {
 	kota := c.DefaultQuery("kota", "all")
	resUsers, err := service.GetBank(c, kota )
	helper.PanicIfError(err)
   
	status := resUsers["status"].(int)
	if status == 200{
		c.JSON(status, gin.H{
		"data": resUsers["data"],
	})
	}else{
		c.JSON(status, gin.H{
		"data": resUsers["data"],
	})
	}
}
func GetBankAntrian(c *gin.Context) {
	bankId := c.Param("id")
	dt := time.Now()
	dateSetting := c.DefaultQuery("date", dt.Format("2006-01-02"))
	fmt.Println(dt.Format("2006-01-02"))
	resUsers, err := service.GetBankAntrian(c, bankId, dateSetting)
	helper.PanicIfError(err)
	c.JSON(200, gin.H{
		"data": resUsers["data"],
	})
}
func BankBooking(c *gin.Context) {
	 var bankBooking model.BankBooking
	 err := c.ShouldBind(&bankBooking)
	 helper.PanicIfError(err)

	resUsers, err := service.BankBooking(c, bankBooking )
	helper.PanicIfError(err)
	c.JSON(200, gin.H{
		"data": resUsers,
	})
}

func BookingBankSelesai(c *gin.Context) {
	BookingBankId := c.Param("bookingBankId")
	resUsers, err := service.BookingBankSelesai(c, BookingBankId)
	helper.PanicIfError(err)
	c.JSON(200, gin.H{
		"data": resUsers,
	})
}

	