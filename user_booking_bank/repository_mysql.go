package userbookingbank

import (
	"context"
	"log"
	"main/config"
	"main/helper"
	"main/models"
)

func GetBooking(ctx context.Context, id string) (models.DefaultResponse, error) {

	var bankBooking models.UserBookingBank

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := "SELECT user_booking_banks.jam_pelayanan, banks.nama_bank, user_booking_banks.keperluan_layanan, user_booking_banks.id, user_booking_banks.tgl_pelayanan  FROM `users` INNER JOIN `user_booking_banks` ON user_booking_banks.user_id = ? INNER JOIN `banks` on banks.id = user_booking_banks.bank_id WHERE user_booking_banks.status != 'done'"

	rowQuery, err := db.QueryContext(ctx, queryText, id)

	if err != nil {
		log.Fatal(err)
	}

	if rowQuery.Next() {

		 err = rowQuery.Scan(
			&bankBooking.JamPelayanan,
			&bankBooking.NamaBank,
			&bankBooking.KeperluanLayanan,
			&bankBooking.Id,
			&bankBooking.TglPelayanan,
		); 
			helper.PanicIfError(err)
	}else{
			return models.DefaultResponse{
				Code:   400,
				Status: "OK",
				Data: map[string]string{
					"message": "error",
				},
			}, nil
		}

	return models.DefaultResponse{
		Code:   200,
		Status: "OK",
		Data:   bankBooking,
	}, nil
}

func PostBooking(ctx context.Context, book models.BookingBank) (models.DefaultResponse, error) {

	db, err := config.MySQL()
	helper.PanicIfError(err)
	
	// SQL := "select username, password from user where username = ? AND password = ?"
	SQL := "INSERT INTO user_booking_banks (user_id, bank_id, jam_pelayanan, tgl_pelayanan, keperluan_layanan, status) VALUES (?,?,?,?,?,?)"
	// rows, err := db.ExecContext(ctx, SQL, 1, 1, "book.JamPelayanan", "book.TglPelayanan", 1, "")
	rows, err := db.ExecContext(ctx, SQL, book.UserId, book.BankId, book.JamPelayanan, book.TglPelayanan, book.KeperluanLayanan, book.Status)
	helper.PanicIfError(err)
	id, err := rows.LastInsertId()
	helper.PanicIfError(err)

	return models.DefaultResponse{
		Code: 200,
		Status: "OK",
		Data: id,
	}, nil
}