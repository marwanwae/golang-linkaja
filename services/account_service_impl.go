package services

import (
	"belajar-go-rest-api/exception"
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/model/domain"
	"belajar-go-rest-api/model/web"
	"belajar-go-rest-api/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground/validator"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
	DB                *sql.DB
	validate          *validator.Validate
}

func NewAccountService(AccountRepository repository.AccountRepository, DB *sql.DB, validate *validator.Validate) AccountService {
	return &AccountServiceImpl{
		AccountRepository: AccountRepository,
		DB:                DB,
		validate:          validate,
	}
}
func (service *AccountServiceImpl) FindyById(ctx context.Context, AccountId int) web.AccountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Account, err2 := service.AccountRepository.FindyById(ctx, tx, AccountId)
	if err2 != nil {
		panic(exception.NewNotFoundError(err2.Error()))
	}

	return helper.ToAccountResponse(Account)
}

func (service *AccountServiceImpl) Transfer(ctx context.Context, account int, request web.TransferRequest) error {

	errorValidate := service.validate.Struct(request)
	helper.PanicIfError(errorValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	FromAccount, err := service.AccountRepository.FindyById(ctx, tx, account)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	ToAccount, err := service.AccountRepository.FindyById(ctx, tx, request.ToAccountNumber)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if FromAccount.Balance < request.Amount {
		err := errors.New("balance is insufficient")
		panic(exception.NewBalanceError(err.Error()))
	}

	FromTransfer := domain.Transfer{
		Account: FromAccount.AccountNumber,
		Amount:  FromAccount.Balance - request.Amount,
	}
	service.AccountRepository.Update(ctx, tx, FromTransfer)

	ToTransfer := domain.Transfer{
		Account: ToAccount.AccountNumber,
		Amount:  ToAccount.Balance + request.Amount,
	}
	service.AccountRepository.Update(ctx, tx, ToTransfer)

	return nil
}
