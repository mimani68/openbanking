package model

import (
	"time"

	"gorm.io/gorm"
)

type CurrencyBase struct {
	gorm.Model
	Id        int       `json:"id" xml:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt"`
}

type Currency struct {
	CurrencyBase

	Title       string `json:"title" xml:"title" gorm:"unique,not null"`
	Symbol      string `json:"symbol" xml:"symbol" gorm:"unique,not null"`
	Digit       int    `json:"digit" xml:"digit" gorm:"not null"`
	Physical    bool   `json:"physical" xml:"physical" default:"true"`
	Description string `json:"description,omitempty" xml:"description"`
	Reference   string `json:"reference,omitempty" xml:"reference"`
}

func (Currency) TableName() string {
	return "app.currency"
}

func (base *Currency) BeforeCreate(tx *gorm.DB) (err error) {
	// base.ID = uuid.New()
	return
}
