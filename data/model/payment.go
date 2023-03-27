package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment struct {
	gorm.Model
	PaymentBase

	Amount int
	Code   string
}

type PaymentMeta struct {
	gorm.Model
	PaymentBase

	PaymentId              int // FK
	RetryNumber            int
	OriginTracingCode      string
	DestinationTracingCode string

	PaymentChannel int // FK
}

type PaymentItem struct {
	gorm.Model
	PaymentBase

	PaymentId int // FK

	Order           int
	Amount          int
	PayeeAccountId  int // FK
	PayeerAccountId int // FK
	Description     string
}

type PaymentCommission struct {
	gorm.Model
	PaymentBase

	PaymentId  int // FK
	Commission int
}

type PaymentResult struct {
	gorm.Model
	PaymentBase

	PaymentId int // FK
	IsSuccess bool

	FailureCode        string
	FailureDescription string
	FailureMessage     string
}

type PaymentChannel struct {
	gorm.Model
	PaymentBase

	Channel string // Satna, Shaba, Paya, Swift
}
