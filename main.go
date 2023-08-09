package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func main() {

	db, err := sql.Open("godror", `user="hr" password="" connectString=""`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()

	const oracleAnonBlock = "BEGIN p(:1, :2); END;"

	stmt, err := db.PrepareContext(ctx, oracleAnonBlock)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	inParam := "hello"
	var outParam string

	stmt.ExecContext(ctx, inParam, sql.Out{Dest: &outParam})

	fmt.Println("Out parameter is", outParam)

}
