package web

type AccountResponse struct {
	AccountNumber int    `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}
