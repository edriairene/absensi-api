package database

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetByNim(t *testing.T) {
	nim := 13515002
	db, err := sql.Open("mysql", "root:@/presence")
	defer db.Close()
	if err != nil {
		t.Fatal("Can't connect to database, %s", err)
	}
	err = GetByNim(nim)
}