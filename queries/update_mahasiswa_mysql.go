package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"nilai-mahasiswa/config"
	"nilai-mahasiswa/models"
)

func UpdateDataMahasiswa(ctx context.Context, mahasiswa models.NilaiMahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET nama='%v', mata_kuliah='%v', indeks_nilai='%v', nilai='%v' WHERE id=%v", table, mahasiswa.Nama, mahasiswa.MataKuliah, mahasiswa.IndeksNilai, mahasiswa.Nilai, id)

	out, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := out.RowsAffected()
	if check == 0 {
		return errors.New("id tidak ditemukan")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
