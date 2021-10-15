package service

import (
	"context"
	"main/config"
	"main/helper"
	"main/model"
)

func GetBank(ctx context.Context, kota string) (map[string]interface{}, error) {

	var banks []model.Bank

	db, err := config.MySQL()
	helper.PanicIfError(err)
	var queryText string
	if kota == "all" {
		queryText = "SELECT id,nama_bank,alamat FROM banks"
		rowQuery, err := db.QueryContext(ctx, queryText)
		// defer rowQuery.Close()
		helper.PanicIfError(err)
		for rowQuery.Next() {
			var bank model.Bank

			if err = rowQuery.Scan(
				&bank.Id,
				&bank.NamaBank,
				&bank.Alamat,
			); err != nil {
				return map[string]interface{}{
			"status"  : 400,
			"data"  : "get data bank failed",
		}, nil
			}

			banks = append(banks, bank)
		}
	} else {
		queryText = "SELECT id,nama_bank,alamat FROM banks WHERE kota=?"
		rowQuery, err := db.QueryContext(ctx, queryText, kota)
		helper.PanicIfError(err)
		for rowQuery.Next() {
			var bank model.Bank

			if err = rowQuery.Scan(
				&bank.Id,
				&bank.NamaBank,
				&bank.Alamat,
			); err != nil {
				return map[string]interface{}{
			"status"  : 400,
			"data"  : "get data bank failed",
		}, nil
			}

			banks = append(banks, bank)
		}
	}

		return map[string]interface{}{
			"status"  : 200,
			"data"  : banks,
		}, nil
}