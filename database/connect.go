package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(db string) *sqlx.DB {
	conn, err := sqlx.Connect("mysql", "root:@(localhost:3306)/"+db)
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}
