package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func logError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Create(conn string) {
	sdb, err := sql.Open("mysql", conn)
	logError(err)

	err = sdb.Ping()
	logError(err)

	db = sdb
}

func Close() {
	err := db.Close()
	logError(err)
}

func CreateInfo() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Info(
			Name VARCHAR(32)
		)
	`)
	logError(err)
}

func InsertInfo(fname string, lname string) {
	res, err := db.Exec(`
		INSERT INTO Info VALUES(
			?
		)
	`, fname+" "+lname)
	logError(err)

	rowCount, err := res.RowsAffected()
	logError(err)

	log.Printf("rows affected: %d", rowCount)
}

func GetName(fullname string, w http.ResponseWriter) {
	row := db.QueryRow(`
		SELECT Name FROM Info
			WHERE Name=?
	`, fullname).Scan()

	if row != sql.ErrNoRows {
		fmt.Fprintf(w, `<p style="color: black; text-align: center;">entry founded</p>`)
	} else {
		fmt.Fprintf(w, `<p style="color: black; text-align: center;">entry not founded</p>`)
	}
}
