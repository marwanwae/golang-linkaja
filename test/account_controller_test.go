package test

import (
	"belajar-go-rest-api/app"
	"belajar-go-rest-api/controller"
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/middleware"
	"belajar-go-rest-api/model/domain"
	"belajar-go-rest-api/repository"
	"belajar-go-rest-api/services"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_linkaja")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	AccountRepository := repository.NewAccountRepository()
	AccountService := services.NewAccountService(AccountRepository, db, validate)
	AccountController := controller.NewAccountController(AccountService)

	router := app.NewRouter(AccountController)

	return middleware.NewAuthMiddleware(router)
}

func truncateAccount(db *sql.DB) {
	db.Exec("DELETE FROM tbl_account;")
	db.Exec("DELETE FROM tbl_customer;")
}

func TestGetAccountSuccess(t *testing.T) {
	db := setupTestDB()
	truncateAccount(db)

	tx, _ := db.Begin()
	AccountRepository := repository.NewAccountRepository()
	AccountRepository.Save(context.Background(), tx, domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		CustomerName:   "budi",
		Balance:        1000000,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/account/555001", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"])
	assert.Equal(t, float64(555001), responseBody["data"].(map[string]interface{})["account_number"])
	assert.Equal(t, "budi", responseBody["data"].(map[string]interface{})["customer_name"])
	assert.Equal(t, float64(1000000), responseBody["data"].(map[string]interface{})["balance"])
}

func TestGetAccountNotFound(t *testing.T) {
	db := setupTestDB()
	truncateAccount(db)

	router := setupRouter(db)

	requestBody := strings.NewReader(``)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/account/555001", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestTransferSuccess(t *testing.T) {
	db := setupTestDB()
	truncateAccount(db)

	tx, _ := db.Begin()
	AccountRepository := repository.NewAccountRepository()
	AccountRepository.Save(context.Background(), tx, domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		CustomerName:   "budi",
		Balance:        1000000,
	})
	AccountRepository.Save(context.Background(), tx, domain.Account{
		AccountNumber:  555002,
		CustomerNumber: 1002,
		CustomerName:   "andi",
		Balance:        1000000,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"to_account_number" : 555002,
		"amount" : 1000
	   }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/api/account/555001/transfer", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "ok", responseBody["status"])
}

func TestTransferAccountNotFound(t *testing.T) {
	db := setupTestDB()
	truncateAccount(db)

	tx, _ := db.Begin()
	AccountRepository := repository.NewAccountRepository()
	AccountRepository.Save(context.Background(), tx, domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		CustomerName:   "budi",
		Balance:        1000000,
	})
	// AccountRepository.Save(context.Background(), tx, domain.Account{
	// 	AccountNumber:  555002,
	// 	CustomerNumber: 1002,
	// 	CustomerName:   "andi",
	// 	Balance:        1000000,
	// })
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"to_account_number" : 555002,
		"amount" : 1000
	   }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/api/account/555001/transfer", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestTransferFailed(t *testing.T) {
	db := setupTestDB()
	truncateAccount(db)

	tx, _ := db.Begin()
	AccountRepository := repository.NewAccountRepository()
	AccountRepository.Save(context.Background(), tx, domain.Account{
		AccountNumber:  555001,
		CustomerNumber: 1001,
		CustomerName:   "budi",
		Balance:        10,
	})
	AccountRepository.Save(context.Background(), tx, domain.Account{
		AccountNumber:  555002,
		CustomerNumber: 1002,
		CustomerName:   "andi",
		Balance:        1000000,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"to_account_number" : 555002,
		"amount" : 1000
	   }`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/api/account/555001/transfer", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "TRANSFER FAILED", responseBody["status"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateAccount(db)

	tx, _ := db.Begin()
	AccountRepository := repository.NewAccountRepository()
	AccountRepository.Save(context.Background(), tx, domain.Account{
		AccountNumber: 123,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/categories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-key", "SALAH")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
	assert.Nil(t, responseBody["data"])
}
