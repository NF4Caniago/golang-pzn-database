package golangpzndatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "golang-pzn:Golang-PZN123@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
