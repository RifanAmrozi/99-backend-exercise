package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
	var err error
	dsn := "user=rifan password=rifan dbname=rental sslmode=disable"
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}
