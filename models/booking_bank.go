package models
type BookingBank struct {
	UserId string         `json:"user_id"`
	BankId string         `json:"bank_id"`
	KeperluanLayanan string `json:"keperluan_layanan"`
	TglPelayanan  string `json:"tgl_pelayanan"`
	JamPelayanan  string `json:"jam_pelayanan"`
	Status string      `json:"status"`
}
