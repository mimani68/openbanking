package dto

import (
	"time"

	"gorm.io/gorm"
)

type CustomerBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Polymorhic mode for customer -> person,institue,business
type Customer struct {
	gorm.Model
	CustomerBase

	Type string `default:"person" json:"type" xml:"type" enum:"person,institue,business"`
}

type CustomerService struct {
	gorm.Model
	CustomerBase

	ServiceId int // FK
	IsActive  bool
	ExpireAt  time.Time
}

type CustomerPrefrence struct {
	gorm.Model
	CustomerBase

	CustomerId int // FK
	ServiceId  int // FK
	Title      string
	Value      string
	File       string
	Type       string `default:"value" json:"type" xml:"type" enum:"file,value"`
	ExpireAt   time.Time
}

type CustomerActivity struct {
	gorm.Model

	CustomerId int // FK
	ServiceId  int // FK
	Title      string
	Value      string
	Time       time.Time
	Result     string
}
