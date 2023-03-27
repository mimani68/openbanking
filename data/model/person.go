package model

import (
	"time"

	"gorm.io/gorm"
)

type PersonBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Person struct {
	gorm.Model
	PersonBase

	Name   string
	Family string
}

type PersonMeta struct {
	gorm.Model

	PersonId     int
	NationalCode string
	Passport     string
	Protratit    string
}

type PersonLegalInquiry struct {
	gorm.Model

	PersonId int
	Title    string
	Value    string
	Issuer   string
	Time     time.Time
}

type PersonStatus struct {
	gorm.Model

	PersonId      int
	CurrectStatus string
	LastStatus    string
	UpdatedTime   time.Time
}

type PersonConnection struct {
	gorm.Model

	PersonId  int
	Title     string
	Value     string
	IsDefualt string
}

type PersonVerificationToken struct {
	gorm.Model

	PersonId int
	Token    string
	ExpireAt time.Time
	Channel  int

	Status string `default:"PENDING" xml:"status" json:"status" enum:"PENDING,VERIFIED"`
}

type PersonVerificationCannel struct {
	gorm.Model

	PersonId int
	Title    string
	Value    string
}
