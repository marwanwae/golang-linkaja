package web

type TransferRequest struct {
	ToAccountNumber int `json:"to_account_number" validate:"required"`
	Amount          int `json:"amount" validate:"required"`
}
