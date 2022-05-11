package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"nilai-mahasiswa/config"
)

func DeleteDataMahasiswa(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id=%v", table, id)

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
