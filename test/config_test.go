package test

import (
	"belajar-go-rest-api/configs"
	"fmt"
	"testing"
)

func init() {
	configs.LoadConfig("../")
}

func TestConfig(t *testing.T) {
	dbUsername := configs.Configs.DB_USERNAME
	dbPassword := configs.Configs.DB_PASSWORD
	dbDatabase := configs.Configs.DB_DATABASE
	dbHost := configs.Configs.DB_HOST
	dbPort := configs.Configs.DB_PORT

	concat := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	fmt.Println(concat)
}
