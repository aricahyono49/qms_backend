package user

import (
	"context"
	"main/config"
	"main/helper"
	"main/models"
)

func Login(ctx context.Context, user models.User) (map[string]string, error) {

	var users models.User
	db, err := config.MySQL()
	helper.PanicIfError(err)
	
	SQL := "select username, password from user where username = ? AND password = ?"
	rows, err := db.QueryContext(ctx, SQL, user.Username ,user.Password)
	helper.PanicIfError(err)

	if rows.Next() {
		err = rows.Scan(&users.Username, &users.Password); 
		helper.PanicIfError(err)
	} 

	if users.Username == ""{
		return map[string]string{
			"status" : "login error",
		}, err
	}
	return map[string]string{
		"username" : users.Username,
		"password" : users.Password,
	}, nil
}