package repository

import (
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type AccountRepositoryImpl struct {
}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

func (repository *AccountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Transfer domain.Transfer) bool {
	SQL := `
		UPDATE tbl_account
		SET balance = ?
		WHERE account_number = ?;`
	_, err := tx.ExecContext(ctx, SQL, Transfer.Amount, Transfer.Account)

	helper.PanicIfError(err)
	return true
}

func (repository *AccountRepositoryImpl) FindyById(ctx context.Context, tx *sql.Tx, AccountNumber int) (domain.Account, error) {
	SQL := `
			SELECT
			c.customer_number, 
			c.name,
			a.account_number,
			a.balance
		FROM
			tbl_customer c
		JOIN
			tbl_account a ON c.customer_number = a.customer_number
		WHERE 
			a.account_number = ?
		LIMIT 1;`

	r, err := tx.QueryContext(ctx, SQL, AccountNumber)
	helper.PanicIfError(err)
	defer r.Close()
	Account := domain.Account{}
	if r.Next() {
		err2 := r.Scan(&Account.CustomerNumber, &Account.CustomerName, &Account.AccountNumber, &Account.Balance)
		helper.PanicIfError(err2)
		return Account, nil
	} else {
		return Account, errors.New("account is not found")
	}
}

func (repository *AccountRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, accounnt domain.Account) domain.Account {
	SQL_cust := `INSERT INTO tbl_customer (customer_number, name)
				VALUES (?, ?);`
	_, err := tx.ExecContext(ctx, SQL_cust, accounnt.CustomerNumber, accounnt.CustomerName)
	helper.PanicIfError(err)

	SQL_Acc := `INSERT INTO tbl_account (account_number, customer_number, balance)
				VALUES (?, ?, ?);`
	_, err = tx.ExecContext(ctx, SQL_Acc, accounnt.AccountNumber, accounnt.CustomerNumber, accounnt.Balance)
	helper.PanicIfError(err)

	return accounnt
}
