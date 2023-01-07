package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectionDatabase() {
	connect, err := sql.Open(Config.SQLDriver, fmt.Sprintf("%s:%s@/%s", Config.DATABASE_USER, Config.DATABASE_PASSWORD, Config.DATABASE))
	if err != nil {
		log.Fatalf("データベース接続に失敗しました。%s", err)
	}
	DB = connect
}

func CloseDatabase() {
	DB.Close()
}
