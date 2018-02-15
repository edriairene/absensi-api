package database

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func InsertPresence(nim int, presence int) (err error) {
	db, err := sql.Open("mysql", "root:@/presence")
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT presence FROM presence WHERE nim = ?", nim)
	if !rows.Next() {
		stmt, _ := db.Prepare("INSERT INTO presence (nim, presence, pres_time) VALUES (?, ?, ?)")
		defer stmt.Close()
		_, err = stmt.Exec(nim, presence, time.Now().Unix()+25200)
		return
	} else {
		stmt, _ := db.Prepare("UPDATE presence SET nim = ?, presence = ?, pres_time = ? WHERE nim = ?")
		defer stmt.Close()
		_, err = stmt.Exec(nim, presence, time.Now().Unix()+25200, nim)
		return
	}
}