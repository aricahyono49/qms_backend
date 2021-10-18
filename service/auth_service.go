package service

import (
	"context"
	"fmt"
	"main/db"
	"main/helper"
	"main/model"
)

func Login(ctx context.Context, user model.User) (map[string]interface{}, error) {

	var users model.User
	db, err := db.MySQL()
	helper.PanicIfError(err)
	
	SQL := "select id, username, password from users where username = ? AND password = ?"
	rows, err := db.QueryContext(ctx, SQL, user.Username ,user.Password)
	helper.PanicIfError(err)

	if rows.Next() {
		err = rows.Scan(&users.Id,&users.Username, &users.Password); 
		helper.PanicIfError(err)
	}else{
		return map[string]interface{}{
			"status"  : 400,
			"message"  : "login failed",
		}, nil
	}

	return map[string]interface{}{
			"status"  : 200,
			"message"  : "login success",
			"user_id"  : users.Id,
		}, nil
}
func Register(ctx context.Context, user model.User) (map[string]interface{}, error) {

	db, err := db.MySQL()
	helper.PanicIfError(err)
	SQL := "INSERT INTO users (username, password) VALUES (?,?)"
	rows, err := db.ExecContext(ctx, SQL, user.Username, user.Password)
	helper.PanicIfError(err)
	fmt.Println(rows);
	return map[string]interface{}{
		"status"  : 200,
		"message"  : "login success",
	}, nil
}