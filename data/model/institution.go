package model

import (
	"time"

	"gorm.io/gorm"
)

type InstitutionBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Institution struct {
	gorm.Model
	InstitutionBase
	Title string
}

type InstitutionMeta struct {
	gorm.Model
	InstituionId int
	Address      string
}

type InstitutionConfig struct {
	gorm.Model
	InstituionId int
	Key          string
	Value        string
	Category     string `json:"category,omitempty" xml:"category"`
	IsActive     bool   `default:"true" json:"isActive" xml:"isActive"`
}

type InstitutionStaff struct {
	gorm.Model
	InstituionId int
	MemeberId    int
	Level        string // admin, master, user, superuser
}
