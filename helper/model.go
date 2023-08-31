package helper

import (
	"belajar-go-rest-api/model/domain"
	"belajar-go-rest-api/model/web"
)

func ToAccountResponse(Account domain.Account) web.AccountResponse {
	return web.AccountResponse{
		AccountNumber: Account.AccountNumber,
		CustomerName:  Account.CustomerName,
		Balance:       Account.Balance,
	}
}

func ToAccountResponses(categories []domain.Account) []web.AccountResponse {
	var AccountResponses []web.AccountResponse
	for _, Account := range categories {
		AccountResponses = append(AccountResponses, ToAccountResponse(Account))
	}

	return AccountResponses
}
