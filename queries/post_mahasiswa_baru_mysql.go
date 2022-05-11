package queries

import (
	"context"
	"fmt"
	"log"
	"nilai-mahasiswa/config"
	"nilai-mahasiswa/models"
)

func InsertMahasiswaBaru(ctx context.Context, mahasiswabaru models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, mata_kuliah, indeks_nilai, nilai) VALUES ('%v', '%v', '%v', %v)", table, mahasiswabaru.Nama, mahasiswabaru.MataKuliah, mahasiswabaru.IndeksNilai, mahasiswabaru.Nilai)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}
	return nil
}
