package bank

import (
	"context"
	"log"
	"main/config"
	"main/helper"
	"main/models"
)

// GetAll
func GetBankAll(ctx context.Context, kota string) (models.DefaultResponse, error) {

	var banks []models.Bank

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	var queryText string
	if kota == ""{
	queryText = "SELECT id,nama_bank,alamat, booking_slot FROM banks"
	rowQuery, err := db.QueryContext(ctx, queryText)
	helper.PanicIfError(err)
	for rowQuery.Next() {
		var bank models.Bank
		
		if err = rowQuery.Scan(
			&bank.IdBank,
			&bank.Name,
			&bank.Address,
			&bank.BookingSlot,
			); 
			err != nil {
			return models.DefaultResponse{
		Code: 200,
		Status: "OK",
		Data: map[string]string{
			"message" : "error",
		},
	}, nil
		}

		banks = append(banks, bank)
	}
	}else{
		queryText = "SELECT id,nama_bank,alamat, booking_slot FROM banks WHERE kota=?"
		rowQuery, err := db.QueryContext(ctx, queryText, kota)
		helper.PanicIfError(err)
		for rowQuery.Next() {
		var bank models.Bank
		
		if err = rowQuery.Scan(
			&bank.IdBank,
			&bank.Name,
			&bank.Address,
			&bank.BookingSlot,
			); 
			err != nil {
			return models.DefaultResponse{
		Code: 200,
		Status: "OK",
		Data: map[string]string{
			"message" : "error",
		},
	}, nil
		}

		banks = append(banks, bank)
	}
	}

	

return models.DefaultResponse{
		Code: 200,
		Status: "OK",
		Data: banks,
	}, nil
}
func GetBankById(ctx context.Context, id string)(models.DefaultResponse, error) {

	var bankDetail models.BankDetail

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := "SELECT user_booking_banks.jam_pelayanan,user_booking_banks.bank_id, user_booking_banks.keperluan_layanan, user_booking_banks.id , banks.nama_bank, banks.booking_slot ,user_booking_banks.tgl_pelayanan, banks.alamat FROM `users` INNER JOIN `user_booking_banks` ON user_booking_banks.user_id = users.id INNER JOIN `banks` on banks.id = user_booking_banks.bank_id WHERE user_booking_banks.status != 'done' AND bank_id = ?;"

	rowQuery, err := db.QueryContext(ctx, queryText, id)

	if err != nil {
		log.Fatal(err)
	}

	if rowQuery.Next() {
		
		 err = rowQuery.Scan(
			&bankDetail.JamPelayanan,
			&bankDetail.BankId,
			&bankDetail.KeperluanLayanan,
			&bankDetail.Id,
			&bankDetail.NamaBank,
			&bankDetail.BookingSlot,
			&bankDetail.TglPelayanan,
			&bankDetail.Alamat,
			); 
			helper.PanicIfError(err)

	}else{
		return models.DefaultResponse{
		Code: 400,
		Status: "NOT OK",
		Data: map[string]string{
			"message" : "error",
		},
	}, nil
	}

return models.DefaultResponse{
		Code: 200,
		Status: "OK",
		Data: bankDetail,
	}, nil
}

