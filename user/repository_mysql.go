package user

import (
	"context"
	"main/config"
	"main/helper"
	"main/models"
)

func Login(ctx context.Context, user models.User) (models.DefaultResponse, error) {

	var users models.User
	db, err := config.MySQL()
	helper.PanicIfError(err)
	
	SQL := "select username, password from users where username = ? AND password = ?"
	rows, err := db.QueryContext(ctx, SQL, user.Username ,user.Password)
	helper.PanicIfError(err)

	if rows.Next() {
		err = rows.Scan(&users.Username, &users.Password); 
		helper.PanicIfError(err)
	}else{
		return models.DefaultResponse{
		Code: 400,
		Status: "login failed",
		Data: nil,
	}, nil
	}

	return models.DefaultResponse{
		Code: 200,
		Status: "OK",
		Data: users,
	}, nil
}
func Register(ctx context.Context, user models.User) (models.DefaultResponse, error) {

	db, err := config.MySQL()
	helper.PanicIfError(err)
	
	// SQL := "select username, password from user where username = ? AND password = ?"
	SQL := "INSERT INTO users (username, password) VALUES (?,?)"
	rows, err := db.ExecContext(ctx, SQL, user.Username ,user.Password)
	helper.PanicIfError(err)
	id, err := rows.LastInsertId()
	helper.PanicIfError(err)

	return models.DefaultResponse{
		Code: 200,
		Status: "OK",
		Data: id,
	}, nil
}