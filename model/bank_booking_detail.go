package model

import "time"

type BankBookingDetail struct {
	Id               int32  `json:"id"`
	NoAntrian        int32  `json:"nomor_antrian"`
	NamaBank         string `json:"nama_bank"`
	Alamat           string `json:"alamat"`
	KeperluanLayanan int32  `json:"keperluan_layanan"`
	JamPelayanan     string `json:"jam_pelayanan"`
	TglPelayanan     time.Time  `json:"tanggal_pelayanan"`
	BookingSLot      int32  `json:"booking_slot"`
}
