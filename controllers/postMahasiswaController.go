package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"nilai-mahasiswa/models"
	"nilai-mahasiswa/queries"
	"nilai-mahasiswa/utils"

	"github.com/julienschmidt/httprouter"
)

func InsertMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mahasiswabaru models.NilaiMahasiswa

	err := json.NewDecoder(r.Body).Decode(&mahasiswabaru)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if mahasiswabaru.Nilai <= 100 && mahasiswabaru.Nilai >= 80 {
		mahasiswabaru.IndeksNilai = "A"
	} else if mahasiswabaru.Nilai >= 70 && mahasiswabaru.Nilai < 80 {
		mahasiswabaru.IndeksNilai = "B"
	} else if mahasiswabaru.Nilai >= 60 && mahasiswabaru.Nilai < 70 {
		mahasiswabaru.IndeksNilai = "C"
	} else if mahasiswabaru.Nilai >= 50 && mahasiswabaru.Nilai < 60 {
		mahasiswabaru.IndeksNilai = "D"
	} else if mahasiswabaru.Nilai < 50 && mahasiswabaru.Nilai >= 0 {
		mahasiswabaru.IndeksNilai = "E"
	}

	if err := queries.InsertMahasiswaBaru(ctx, mahasiswabaru); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Input Success",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}
