package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	byteData, err := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "error response", http.StatusBadRequest)
	}

	w.WriteHeader(status)
	w.Write([]byte(byteData))
}
