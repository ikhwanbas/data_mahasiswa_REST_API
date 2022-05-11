package models

import "time"

type NilaiMahasiswa struct {
	ID          int       `json:"id"`
	Nama        string    `json:"nama"`
	MataKuliah  string    `json:"mata-kuliah"`
	IndeksNilai string    `json:"indeks-nilai"`
	Nilai       int       `json:"nilai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
