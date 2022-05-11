package controllers

import (
	"context"
	"fmt"
	"net/http"
	"nilai-mahasiswa/queries"
	"nilai-mahasiswa/utils"

	"github.com/julienschmidt/httprouter"
)

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMahasiswa = ps.ByName("id")

	if err := queries.DeleteDataMahasiswa(ctx, idMahasiswa); err != nil {
		message := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, message, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Delete Success",
	}
	utils.ResponseJSON(w, res, http.StatusOK)

}
