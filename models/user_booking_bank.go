package models

type UserBookingBank struct {
	Id string `json:"id"`
	NamaBank string `json:"nama_bank"`
	KeperluanLayanan string `json:"keperluan_layanan"`
	JamPelayanan string `json:"jam_pelayanan"`
	TglPelayanan string `json:"tgl_pelayanan"`
}
