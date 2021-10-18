package service

import (
	"context"
	// "fmt"
	"main/db"
	"main/helper"
	"main/model"
	// "time"
)

func GetUserBookingBank(ctx context.Context, user_id string) (map[string]interface{}, error) {
	var booking_bank_detail model.BankBookingDetail
	db, err := db.MySQL()
	helper.PanicIfError(err)
	
	queryText := "SELECT user_booking_banks.id, user_booking_banks.nomor_antrian, banks.nama_bank, banks.alamat, user_booking_banks.keperluan_layanan, user_booking_banks.jam_pelayanan, user_booking_banks.tanggal_pelayanan, banks.booking_slot  FROM `users` INNER JOIN `user_booking_banks` ON user_booking_banks.user_id = users.id INNER JOIN `banks` on banks.id = user_booking_banks.bank_id WHERE user_booking_banks.status != 'done' AND user_booking_banks.user_id = ?"
	rowQuery, err := db.QueryContext(ctx, queryText, user_id)
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
		"data": nil,
	}, nil
	}
	defer rowQuery.Close()


// input := "2017-08-31"
// layout := "2006-01-02"
// t, _ := time.Parse(layout, input)
// fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
// fmt.Println(t.Format("02-Jan-2006"))
	
	return map[string]interface{}{
		"data": booking_bank_detail,
	}, nil
}