package dto

import "time"

type AccountBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Account struct {
	AccountBase
	Customer int // FK

	AcceptablePaymentMethod string

	Bank        string
	BankCode    string
	BankId      int
	AccountCode string

	Balance int
}
