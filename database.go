package golangpzndatabase

import (
	"database/sql"
	"fmt"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	fmt.Println("err database.go")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
