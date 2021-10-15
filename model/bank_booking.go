package model

type BankBooking struct {
	UserId int32 `json:"user_id"`
	BankId int32 `json:"bank_id"`
	KeperluanLayanan int32 `json:"keperluan_layanan"`
	NamaBank string `json:"nama_bank"`
}