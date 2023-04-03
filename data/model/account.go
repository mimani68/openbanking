package model

import (
	"time"

	"gorm.io/gorm"
)

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

func (AccountBase) TableName() string {
	return "app.account"
}

func (base *AccountBase) BeforeCreate(tx *gorm.DB) (err error) {
	// base.ID = uuid.New()
	return
}
