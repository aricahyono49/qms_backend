package models

type Bank struct {
	IdBank string `json:"id"`
	Name string `json:"nama_bank"`
	Address string `json:"alamat"`
	BookingSlot int `json:"booking_slot"`
}
