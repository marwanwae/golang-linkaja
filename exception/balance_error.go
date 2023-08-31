package exception

type BalanceError struct {
	Error string
}

func NewBalanceError(error string) BalanceError {
	return BalanceError{Error: error}
}
