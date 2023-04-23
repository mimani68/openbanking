package model

import (
	"time"

	"gorm.io/gorm"
)

type PersonBase struct {
	gorm.Model
	Id        int       `json:"id" xml:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" xml:"deletedAt"`
}

type Person struct {
	PersonBase

	Name   string `json:"name" xml:"name"`
	Family string `json:"family" xml:"family"`

	Customer Customer `gorm:"polymorphic:Ref;polymorphicValue:person"`
}

type PersonMeta struct {
	gorm.Model

	PersonId int    `json:"personId" xml:"personId"`
	Person   Person `gorm:"foreignKey:PersonId"`

	NationalCode string `json:"nationalCode" xml:"nationalCode" gorm:"index,not null"`
	Passport     string `json:"passport" xml:"passport" gorm:"index,not null"`
	Portrait     string `json:"Portrait" xml:"Portrait"`
}

type PersonLegalInquiry struct {
	gorm.Model

	PersonId int    `json:"personId" xml:"personId"`
	Person   Person `gorm:"foreignKey:PersonId"`

	Title  string    `json:"title" xml:"title" gorm:"index,not null"`
	Value  string    `json:"value" xml:"value" gorm:"index,not null"`
	Issuer string    `json:"issuer" xml:"issuer" gorm:"index,not null"`
	Time   time.Time `json:"time" xml:"time" gorm:"index,not null"`
}

type PersonStatus struct {
	gorm.Model

	PersonId int    `json:"personId" xml:"personId"`
	Person   Person `gorm:"foreignKey:PersonId"`

	CurrentStatus string    `json:"currentStatus" xml:"currentStatus"`
	LastStatus    string    `json:"lastStatus" xml:"lastStatus"`
	UpdatedTime   time.Time `json:"updatedTime" xml:"updatedTime"`
}

type PersonConnectionChannel struct {
	gorm.Model

	PersonId int    `json:"personId" xml:"personId"`
	Person   Person `gorm:"foreignKey:PersonId"`

	Title      string `json:"title" xml:"title" gorm:"index,not null"`
	Value      string `json:"value" xml:"value" gorm:"index,not null"`
	IsVerified bool   `json:"isVerified" xml:"isVerified" default:"false"`
	Order      int    `json:"order" xml:"order"`
	IsDefault  string `json:"isDefault" xml:"isDefault"`
}

type PersonVerificationToken struct {
	gorm.Model

	PersonId int    `json:"personId" xml:"personId"`
	Person   Person `gorm:"foreignKey:PersonId"`

	Token    string    `json:"token" xml:"token"`
	ExpireAt time.Time `json:"expireAt" xml:"expireAt"`

	ChannelId int `json:"channelId" xml:"channelId"`
	Channel   int `gorm:"foreignKey:ChannelId"`

	Status string `default:"PENDING" xml:"status" json:"status" enum:"PENDING,VERIFIED"`
}

type PersonVerificationCannel struct {
	gorm.Model

	PersonId int    `json:"personId" xml:"personId"`
	Person   Person `gorm:"foreignKey:PersonId"`

	Title string `json:"title" xml:"title" gorm:"index,not null"`
	Value string `json:"value" xml:"value" gorm:"index,not null"`
}
