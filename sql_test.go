package golangpzndatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("err sql_test")
	fmt.Println("Success insert new customer")
}
