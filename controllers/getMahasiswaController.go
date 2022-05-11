package controllers

import (
	"context"
	"log"
	"net/http"
	"nilai-mahasiswa/queries"
	"nilai-mahasiswa/utils"

	"github.com/julienschmidt/httprouter"
)

func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	allmahasiswa, err := queries.GetAllMahasiswa(ctx)
	if err != nil {
		log.Fatal(err)
	}

	utils.ResponseJSON(w, allmahasiswa, http.StatusOK)
}
