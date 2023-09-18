package belajargolangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	//connection pooling
	db.SetMaxIdleConns(10)                  // minmal koneksi
	db.SetMaxOpenConns(100)                 // maximal koneksi
	db.SetConnMaxIdleTime(5 * time.Minute)  // jika selama 5 menit bengong makan koneksi di hapus
	db.SetConnMaxLifetime(60 * time.Minute) // lama penggunaan koneksi jika sudah 60 emnit maka akan diganti dg koneksi yang baru

	return db
}
