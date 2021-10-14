package models

type BankDetail struct {
	Id string `json:"id"`
	BankId string `json:"bank_id"`
	NamaBank string `json:"nama_bank"`
	KeperluanLayanan string `json:"keperluan_layanan"`
	JamPelayanan string `json:"jam_pelayanan"`
	TglPelayanan string `json:"tgl_pelayanan"`
	BookingSlot string `json:"booking_slot"`
	Alamat string `json:"alamat"`
}