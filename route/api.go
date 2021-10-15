package route

import (
	"main/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// User
	r.POST("/api/user/login", repository.Login)
	r.POST("/api/user/register", repository.Register)
	r.GET("/api/user/:id/booking", repository.GetUserBookingBank)

	// Post 
	r.GET("/api/bank", repository.BankIndex)
	r.POST("/api/bank/booking", repository.BankBooking)
	r.DELETE("/api/bank/booking/:bookingBankId", repository.BookingBankSelesai)
	r.GET("/api/bank/:id/antrian", repository.GetBankAntrian)
	return r
}