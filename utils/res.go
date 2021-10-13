package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)
w.Header().Set("Access-Control-Allow-Origin", "*")

// semua method diperbolehkan masuk
w.Header().Set("Access-Control-Allow-Methods", "*")

// semua header diperbolehkan untuk disisipkan
w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "error om", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	if status == 200{
		w.WriteHeader(status)
	}
	w.Write([]byte(ubahkeByte))
}