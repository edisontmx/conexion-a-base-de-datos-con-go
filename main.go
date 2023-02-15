package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ctx := context.Background()
	db, err := createCon()
	if err != nil {
		panic(err)
	}

	/* err = insert(ctx, db, "el hombre mas rico de babilonia", "George Samuel Clason", time.Time{2020/11/})
	if err != nil {
		panic(err)
	} */
	type libro struct {
		id     int64
		nombre string
		autor  string
		fecha  time.Time
	}

	err = buscarLibros(ctx, db)
	if err != nil {
		panic(err)
	}

	db.Close()

}

func createCon() (*sql.DB, error) {
	conexion := "root:@tcp(localhost:3306)/sistema?parseTime=True"

	db, err := sql.Open("mysql", conexion)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func buscarLibros(ctx context.Context, db *sql.DB) error {
	qry := `SELECT id, nombre, autor, fecha  FROM  libros  WHERE  id =?`
	row := db.QueryRowContext(ctx, qry, 4)

	var id int64
	var nombre string
	var autor string
	var fecha time.Time

	err := row.Scan(&id, &nombre, &autor, &fecha)
	if err != nil {
		return err
	}
	log.Println("ROW:", id, nombre, autor, fecha)
	return nil
}

func insert(ctx context.Context, db *sql.DB, nombre string, autor string, fecha time.Time) error {

	qryAdd := `INSERT INTO libros( nombre, autor, fecha) values(?,?,?) `

	result, err := db.ExecContext(ctx, qryAdd, nombre, autor, fecha)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("INSERTED ID: ", id)

	return nil
}
