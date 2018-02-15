package database

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type PresModel struct {
	PresTime int64
}

func InsertPresence(nim int, presence int) (err error) {

	presModel := PresModel{}
	count := 0

	db, err := sql.Open("mysql", "root:@/presence")
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT max(pres_time) FROM presence WHERE nim = ?", nim)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		count++
		err := rows.Scan(&presModel.PresTime)
		if err != nil {
			///log.Fatal(err)
		}
	}

	if count == 1 {
		var timeNow = time.Now().Unix()+25200
		var dayMax = presModel.PresTime/86400
		var dayNow = timeNow/86400

		if(dayNow-dayMax > 0) {
			stmt, _ := db.Prepare("INSERT INTO presence (nim, presence, pres_time) VALUES (?, ?, ?)")
			defer stmt.Close()
			_, err = stmt.Exec(nim, presence, time.Now().Unix()+25200)
		} else {
			stmt, _ := db.Prepare("UPDATE presence SET nim = ?, presence = ?, pres_time = ? WHERE nim = ?")
			defer stmt.Close()
			_, err = stmt.Exec(nim, presence, time.Now().Unix()+25200, nim)
		}
	} else {
		stmt, _ := db.Prepare("INSERT INTO presence (nim, presence, pres_time) VALUES (?, ?, ?)")
		defer stmt.Close()
		_, err = stmt.Exec(nim, presence, time.Now().Unix()+25200)
	}
	return
}