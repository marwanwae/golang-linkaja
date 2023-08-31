package services

import (
	"belajar-go-rest-api/model/web"
	"context"
)

type AccountService interface {
	Transfer(ctx context.Context, account int, request web.TransferRequest) error
	FindyById(ctx context.Context, AccountId int) web.AccountResponse
}
