package database

import (
	"database/sql"
	"time"
)

func InsertPresence(nim int, presence int) (err error) {
	db, err := sql.Open("mysql", "root:@/presence")
	if err != nil {
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO presence (nim, presence, pres_time) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(nim, presence, time.Now().Unix())
	return
}