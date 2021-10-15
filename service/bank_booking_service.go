package service

import (
	"context"
	"fmt"
	"main/config"
	"main/helper"
	"main/model"
	"time"
)


func GetBankAntrian(ctx context.Context, bankId string) (map[string]interface{}, error) {
	var booking_bank_detail model.BankBookingDetail
	db, err := config.MySQL()
	helper.PanicIfError(err)
	queryText := "SELECT user_booking_banks.id, user_booking_banks.nomor_antrian, banks.nama_bank, banks.alamat, user_booking_banks.keperluan_layanan, user_booking_banks.jam_pelayanan, user_booking_banks.tanggal_pelayanan, banks.booking_slot  FROM `users` INNER JOIN `user_booking_banks` ON user_booking_banks.user_id = users.id INNER JOIN `banks` on banks.id = user_booking_banks.bank_id WHERE user_booking_banks.status != 'done' AND user_booking_banks.bank_id = ?"
	rowQuery, err := db.QueryContext(ctx, queryText, bankId)
	helper.PanicIfError(err)
	if rowQuery.Next(){
		rowQuery.Scan(
			&booking_bank_detail.Id,
			&booking_bank_detail.NoAntrian,
					&booking_bank_detail.NamaBank,
					&booking_bank_detail.Alamat,
				&booking_bank_detail.KeperluanLayanan,
			&booking_bank_detail.JamPelayanan,
			&booking_bank_detail.TglPelayanan,
		&booking_bank_detail.BookingSLot )

	}else{
		return map[string]interface{}{
		"status": 200,
		"data":   nil,
	}, nil
	}
	defer rowQuery.Close()
	
	return map[string]interface{}{
		"status": 200,
		"data":   booking_bank_detail,
	}, nil
}

func BankBooking(ctx context.Context, book model.BankBooking) (map[string]interface{}, error) {

	db, err := config.MySQL()
	helper.PanicIfError(err)


	SQL3 := "SELECT booking_slot FROM banks WHERE id = ? "
	checkBookingSlot , err := db.QueryContext(ctx, SQL3, book.BankId)
	helper.PanicIfError(err)
	var bookingSLot int32
	const kuota int32 = 20
	if checkBookingSlot.Next(){
		checkBookingSlot.Scan(&bookingSLot)
	}

	tx, err := db.Begin(); 
	helper.PanicIfError(err)
	dt := time.Now()

	SQL := "INSERT INTO user_booking_banks (user_id, bank_id, jam_pelayanan, tanggal_pelayanan, keperluan_layanan, nomor_antrian) VALUES (?,?,?,?,?,?)"
	rows , err := tx.ExecContext(ctx, SQL, book.UserId, book.BankId, "08:00", dt, book.KeperluanLayanan, (kuota - (bookingSLot - 1)) )
	helper.PanicIfError(err)
	SQL2 := "UPDATE banks SET banks.booking_slot = (SELECT banks.booking_slot FROM banks WHERE banks.id = ? ) - 1 WHERE banks.id = ?"
	rowku , err := tx.ExecContext(ctx, SQL2, book.BankId, book.BankId)
	helper.PanicIfError(err)
	fmt.Println(rowku);
	fmt.Println(rows);
	
	if bookingSLot <= 0{
		tx.Rollback()
	}else{
		tx.Commit()
	}
	helper.PanicIfError(err)
	return map[string]interface{}{
		"kuota"	  	   : bookingSLot,
		"nomor_antrian": (kuota - (bookingSLot - 1)),
		"nama_bank"    : book.NamaBank,
	}, nil
}

func BookingBankSelesai(ctx context.Context, bookingBankId string) (string, error) {
	db, err := config.MySQL()
	helper.PanicIfError(err)
	queryText := "UPDATE user_booking_banks SET status= ? WHERE user_booking_banks.id = ?"
	rowQuery, err := db.ExecContext(ctx, queryText, "done", bookingBankId )
	helper.PanicIfError(err)
	fmt.Println(rowQuery);
	
	return "delete layanan success", nil
}
