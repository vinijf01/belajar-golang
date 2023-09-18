package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko')"
	script := "UPDATE customer SET name = 'Joko Baru' WHERE id = 'joko'"
	// script := "DELETE FROM customer WHERE id = 'joko'"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("name", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, nama, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, nama string
		var email sql.NullString //untuk yang nilainya bisa null
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool
		err = rows.Scan(&id, &nama, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("======================")
		fmt.Println("id", id)
		fmt.Println("name", nama)
		if email.Valid {
			fmt.Println("email", email.String)

		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		if birth_date.Valid {
			fmt.Println("birth_date", birth_date.Time.String())

		}
		fmt.Println("married", married)
		fmt.Println("created_at", created_at)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	//string concat // script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1 "
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1 "
	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "vini"
	password := "vini"
	// script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko')"
	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	// script := "UPDATE customer SET name = 'Joko Baru' WHERE id = 'joko'"
	// script := "DELETE FROM customer WHERE id = 'joko'"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "vini@gmail.com"
	comment := "Test Comment"
	// script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko')"
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// script := "UPDATE customer SET name = 'Joko Baru' WHERE id = 'joko'"
	// script := "DELETE FROM customer WHERE id = 'joko'"
	result, err := db.ExecContext(ctx, script, username, comment)
	if err != nil {
		panic(err)
	}
	inserId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", inserId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	//use koneksi yang sama jika pake prepare statement
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Koment ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	//do transaction
	for i := 0; i < 10; i++ {
		email := "txeko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Dari TxKoment ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
	err = tx.Rollback() //rollback untuk membatalkan , data tidak jadi ditambahakan ke db, kalo ingin add pake commit
	if err != nil {
		panic(err)
	}
}
