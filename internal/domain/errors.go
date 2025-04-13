package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrDuplicatedAPIKey = errors.New("api key already existis")
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrUnauthorizedAcess = errors.New("unauthorized access")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
)	