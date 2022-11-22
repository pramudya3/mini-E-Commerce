package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetCoonectMysql() *sql.DB {
	db, err := sql.Open("mysql", "lemper:lemper@tcp(localhost:3306)/ecommerce?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
