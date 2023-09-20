package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_db_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_db_migration" -path db/migrations up
//migrate create -ext sql -dir db/migrations create_table_first
// migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_db_migration" -path db/migrations version
// migrate -database "mysql://root:@tcp(localhost:3306)/belajar_golang_db_migration" -path db/migrations force 20230919133520
