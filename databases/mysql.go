package databases

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetCoonectMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/e_commerce?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
