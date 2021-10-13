package main

import (
	"context"
	"encoding/json"
	"fmt"
	"main/config"
	"main/helper"
	"main/models"
	"main/user"
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