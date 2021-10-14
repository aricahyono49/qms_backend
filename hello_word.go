package main

import (
	"context"
	"encoding/json"
	"fmt"
	"main/bank"
	"main/config"
	"main/helper"
	"main/models"
	"main/user"
	"main/user_booking_bank"
	"main/utils"
	"net/http"
)

func main() {

	db, e := config.MySQL()
	helper.PanicIfError(e)
	eb := db.Ping()
	helper.PanicIfError(eb)
	fmt.Println("database running success")

	http.HandleFunc("/api/login", Login)
	http.HandleFunc("/api/register", Register)
	http.HandleFunc("/api/bank", GetBankAll)
	http.HandleFunc("/api/bank/detail", GetBankById)
	http.HandleFunc("/api/book", GetBooking)
	http.HandleFunc("/api/book/post", PostBooking)
	
	err := http.ListenAndServe(":9000", nil)
	helper.PanicIfError(err)

	
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.ResponseJSON(w, "resUsers", http.StatusOK)
	}
	if r.Method == "POST" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var users models.User
		err := json.NewDecoder(r.Body).Decode(&users);
		helper.PanicIfError(err)
		
		resUsers, err := user.Login(ctx, users)
		helper.PanicIfError(err)

		utils.ResponseJSON(w, resUsers, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
}

// Login
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.ResponseJSON(w, "resUsers", http.StatusOK)
	}
	if r.Method == "POST" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var users models.User
		err := json.NewDecoder(r.Body).Decode(&users);
		helper.PanicIfError(err)
		
		resUsers, err := user.Register(ctx, users)
		helper.PanicIfError(err)

		utils.ResponseJSON(w, resUsers, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
}

// GetMahasiswa
func GetBankAll(w http.ResponseWriter, r *http.Request) {



	if r.Method == "OPTIONS" {
		utils.ResponseJSON(w, "resUsers", http.StatusOK)
	}
var kotaBank string = r.URL.Query().Get("kota")
	 if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

	banks, err := bank.GetBankAll(ctx, kotaBank)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, banks, http.StatusOK)
		return
	}


	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}
// GetMahasiswa
func GetBankById(w http.ResponseWriter, r *http.Request) {
var idBank string = r.URL.Query().Get("id")

	if r.Method == "OPTIONS" {
		utils.ResponseJSON(w, "resUsers", http.StatusOK)
	}
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

	bank, err := bank.GetBankById(ctx,idBank)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, bank, http.StatusOK)
		return
	}
	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return

}
func GetBooking(w http.ResponseWriter, r *http.Request) {
var idBank string = r.URL.Query().Get("id")

	if r.Method == "OPTIONS" {
		utils.ResponseJSON(w, "resUsers", http.StatusOK)
	}
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

	bank, err := userbookingbank.GetBooking(ctx,idBank)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, bank, http.StatusOK)
		return
	}
	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return

}


func PostBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		utils.ResponseJSON(w, "resUsers", http.StatusOK)
	}
	if r.Method == "POST" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var users models.BookingBank
		err := json.NewDecoder(r.Body).Decode(&users);
		helper.PanicIfError(err)
		
		resUsers, err := userbookingbank.PostBooking(ctx, users)
		helper.PanicIfError(err)

		utils.ResponseJSON(w, resUsers, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
}


