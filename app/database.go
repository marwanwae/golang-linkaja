package app

import (
	"belajar-go-rest-api/configs"
	"belajar-go-rest-api/helper"
	"database/sql"
	"fmt"
	"time"
)

func NewDB() *sql.DB {
	dbUsername := configs.Configs.DB_USERNAME
	dbPassword := configs.Configs.DB_PASSWORD
	dbDatabase := configs.Configs.DB_DATABASE
	dbHost := configs.Configs.DB_HOST
	dbPort := configs.Configs.DB_PORT

	concat := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	db, err := sql.Open("mysql", concat)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
