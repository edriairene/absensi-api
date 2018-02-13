package database

import (
	"testing"
	"database/sql"
)

func TestInsertPresence(t *testing.T) {
	nim := 13515999
	presence := 1
	err := InsertPresence(nim, presence)
	if err != nil {
		t.Fatal(err)
	}
	db, err := sql.Open("mysql", "root:@/presence")
	defer db.Close()
	if err != nil {
		t.Fatal("Cannot connect to database")
	}
	rows, err := db.Query("SELECT presence FROM presence WHERE nim = ?", nim)
	if !rows.Next() {
		t.Fatalf("Want %s, got zero result", nim)
	}
}