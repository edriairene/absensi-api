package database

import (
	"absensi-api/database/models"
	"database/sql"
	"log"
)

func GetByNim(nim int) (err error) {

	student := models.Student{}
	db, err := sql.Open("mysql", "root:@/presence")
	defer db.Close()
	if err != nil {
		return
	}
	rows, err := db.Query("SELECT * FROM student WHERE nim = ?", nim)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&student.Nim, &student.Password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(student)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}