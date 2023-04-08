package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerBase struct {
	gorm.Model
	Id        int       `json:"id" xml:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" xml:"deletedAt"`
}

// Polymorhic mode for customer -> person,institue,business
type Customer struct {
	CustomerBase

	Title string `json:"title" xml:"title"`
	Type  string `json:"type" xml:"type" enum:"person,institue,business"`
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

	Title    string
	Value    string
	File     string
	Type     string `json:"type" xml:"type" enum:"file,value" default:"value"`
	ExpireAt time.Time

	CustomerId int // FK
	ServiceId  int // FK
}

type CustomerActivity struct {
	gorm.Model

	Title  string `json:"title" xml:"title" enum:"opening,closing,suspending,under-investigation" default:"opening"`
	Value  string
	Time   time.Time
	Result string

	CustomerId int // FK
	ServiceId  int // FK
}
