package model

import (
	"time"

	"gorm.io/gorm"
)

type ServiceBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service struct {
	gorm.Model
	ServiceBase

	Title       string
	Description string
}

type ServiceStatus struct {
	gorm.Model
	ServiceBase

	ServiceId int
	Status    string
}
