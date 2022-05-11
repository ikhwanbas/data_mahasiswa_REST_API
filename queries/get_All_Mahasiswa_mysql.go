package queries

import (
	"context"
	"fmt"
	"log"
	"nilai-mahasiswa/config"
	"nilai-mahasiswa/models"
	"time"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAllMahasiswa(ctx context.Context) ([]models.NilaiMahasiswa, error) {
	var allmahasiswa []models.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't Connect To MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mahasiswa models.NilaiMahasiswa
		var createdAt, updatedAt string

		if err = rowQuery.Scan(
			&mahasiswa.ID,
			&mahasiswa.Nama,
			&mahasiswa.MataKuliah,
			&mahasiswa.IndeksNilai,
			&mahasiswa.Nilai,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		// change format string to datetime for created_at and updated_at
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		allmahasiswa = append(allmahasiswa, mahasiswa)

	}

	return allmahasiswa, nil

}
