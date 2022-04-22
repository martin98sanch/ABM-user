package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type (
	ExecFunc  func(query string, args ...interface{}) (sql.Result, error)
	QueryFunc func(query string, args ...interface{}) (sql.Rows, error)
)

const url = "root:123456@tcp(localhost:3306)/goweb_db"

var db *sql.DB

//Realizar lac conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection
}

//Cerrar la Conexion
func Close() {
	db.Close()
}

//Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	defer Close()
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

//Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	defer Close()
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
