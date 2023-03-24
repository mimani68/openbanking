package dto

import (
	"time"

	"gorm.io/gorm"
)

type PersonBaseDto struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PersonDto struct {
	gorm.Model
	PersonBaseDto
	DestinationBankCode    string
	DestinationBankId      int
	DestinationAccountCode string
	DestinationAccountId   int
}

type PersonStatusDto struct {
	gorm.Model
	PersonId      int
	CurrectStatus string
	LastStatus    string
	UpdatedTime   time.Time
}

type PersonConnectionDto struct {
	gorm.Model
	PersonId int
}

type PersonVerificationTokenDto struct {
	gorm.Model
	PersonId int
}

type PersonVerificationCannelDto struct {
	gorm.Model
	PersonId int
}

// db.Model(&PersonDto{}).HasMany(&PersonStatusDto{})
