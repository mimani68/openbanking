package dto

import (
	"time"

	"gorm.io/gorm"
)

type InstitutionBaseDto struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InstitutionDto struct {
	gorm.Model
	InstitutionBaseDto
	Title string
}

type InstitutionConfigDto struct {
	gorm.Model
	InstituionId int
	Key          string
	Value        string
	Category     string `json:"category,omitempty" xml:"category"`
	IsActive     bool   `default:"true" json:"isActive" xml:"isActive"`
}

type InstitutionStaffDto struct {
	gorm.Model
	InstituionId int
	MemeberId    int
	Level        string // admin, master, user, superuser
}
