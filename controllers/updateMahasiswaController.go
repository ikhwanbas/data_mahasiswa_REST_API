package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"nilai-mahasiswa/models"
	"nilai-mahasiswa/queries"
	"nilai-mahasiswa/utils"

	"github.com/julienschmidt/httprouter"
)

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mahasiswa models.NilaiMahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if mahasiswa.Nilai <= 100 && mahasiswa.Nilai >= 80 {
		mahasiswa.IndeksNilai = "A"
	} else if mahasiswa.Nilai >= 70 && mahasiswa.Nilai < 80 {
		mahasiswa.IndeksNilai = "B"
	} else if mahasiswa.Nilai >= 60 && mahasiswa.Nilai < 70 {
		mahasiswa.IndeksNilai = "C"
	} else if mahasiswa.Nilai >= 50 && mahasiswa.Nilai < 60 {
		mahasiswa.IndeksNilai = "D"
	} else if mahasiswa.Nilai < 50 && mahasiswa.Nilai >= 0 {
		mahasiswa.IndeksNilai = "E"
	} else {
		utils.ResponseJSON(w, "nilai tidak boleh lebih dari 100", http.StatusBadRequest)
		return
	}

	var idMahasiswa = ps.ByName("id")

	if err := queries.UpdateDataMahasiswa(ctx, mahasiswa, idMahasiswa); err != nil {
		message := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, message, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "successfully updated",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)

}
