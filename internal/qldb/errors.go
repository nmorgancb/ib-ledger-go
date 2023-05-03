package qldb

import "fmt"

type AccountNotFoundError struct {
	Id string
}

func (e *AccountNotFoundError) Error() string {
	return fmt.Sprintf("account not found: %s", e.Id)
}

type InsufficientBalanceError struct {
}

func (e *InsufficientBalanceError) Error() string {
	return "insufficient available balance"
}
