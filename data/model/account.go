package model

import (
	"time"

	"gorm.io/gorm"
)

type AccountBase struct {
	gorm.Model
	Id        int       `json:"id" xml:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt" gorm:"index,not null"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt" gorm:"index,not null"`
}

type Account struct {
	gorm.Model
	AccountBase

	AccountCode string

	Customer   Customer `json:"customer,omitempty" xml:"customer" gorm:"foreignKey:CustomerId"`
	CustomerId int      `json:"customerId" xml:"customerId" gorm:"not null"`

	PaymentChannels  []PaymentChannel `json:"paymentChannel,omitempty" xml:"paymentChannel" gorm:"foreignKey:PaymentChannelId"`
	PaymentChannelId []int            `json:"paymentChannelId" xml:"PaymentChannelId"`

	Bank   Bank `json:"bank" xml:"bank" gorm:"foreignKey:BankId"`
	BankId int  `json:"bankId" xml:"bankId"`

	Currency   Currency `gorm:"foreignKey:CurrencyId"`
	CurrencyId int      `json:"currencyId" xml:"currencyId"`

	Balance int `json:"balance" xml:"balance"`
}

func (Account) TableName() string {
	return "app.account"
}

func (base *Account) BeforeCreate(tx *gorm.DB) (err error) {
	// base.ID = uuid.New()
	return
}

type AccountType struct {
	AccountBase

	Type string `json:"accountType" xml:"accountType"`
}

func (AccountType) TableName() string {
	return "app.account_type"
}
