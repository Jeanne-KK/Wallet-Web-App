package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"log"
)

var DB *sql.DB

func Init() error{
	var err error
	DB, err = sql.Open("mysql", "root:1234@tcp(db:3306)/mydatabase")
	if err != nil{
		return fmt.Errorf("fail to open database: %w", err)
	}

	err = DB.Ping()
	if err != nil{
		return fmt.Errorf("fail to ping database: %w", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(10 * time.Minute)

	log.Println("connected database success")
	return nil
}