package repository

import (
	"belajar-go-rest-api/model/domain"
	"context"
	"database/sql"
)

type AccountRepository interface {
	Update(ctx context.Context, tx *sql.Tx, Account domain.Transfer) bool
	FindyById(ctx context.Context, tx *sql.Tx, AccountId int) (domain.Account, error)
	Save(ctx context.Context, tx *sql.Tx, category domain.Account) domain.Account
}
