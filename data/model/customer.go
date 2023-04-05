package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerBase struct {
	gorm.Model
	Id        int       `json:"id" xml:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt" gorm:"index,not null"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt" gorm:"index,not null"`
}

// Polymorhic mode for customer -> person,institue,business
type Customer struct {
	CustomerBase

	Type string `default:"person" json:"type" xml:"type" enum:"person,institue,business"`
}

type CustomerService struct {
	CustomerBase

	Service   Service `gorm:"foreignKey:ServiceId"`
	ServiceId int     `json:"serviceId" xml:"serviceId"`

	IsActive bool      `json:"isActive" xml:"isActive" default:"false"`
	ExpireAt time.Time `json:"expireAt" xml:"expireAt"`
}

type CustomerPrefrence struct {
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
