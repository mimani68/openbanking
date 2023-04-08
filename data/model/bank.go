package model

import (
	"time"

	"gorm.io/gorm"
)

type BankBase struct {
	gorm.Model
	Id        int       `json:"id" xml:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt"`
}

type Bank struct {
	BankBase

	Title   string `json:"title" xml:"title" gorm:"index"`
	Code    string `json:"code" xml:"code" gorm:"index"`
	Address string `json:"address" xml:"address"`
}

func (Bank) TableName() string {
	return "app.bank"
}
