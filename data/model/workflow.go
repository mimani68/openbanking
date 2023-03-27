package model

import (
	"time"

	"gorm.io/gorm"
)

type WorkFlowBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time

	Status string `default:"PENDING" json:"status" xml:"status" enum:"DONE,PENDING"`
}

type WorkFlowAbstract struct {
	gorm.Model
	WorkFlowBase

	Title       string
	Description string

	BPMNFile   string
	HasDiagram bool `default:"false"`
}

type WorkFlow struct {
	gorm.Model
	WorkFlowBase

	WorkFlowAbstractId int
	PaymentId          int // FK
}

type WorkFlowStep struct {
	gorm.Model
	WorkFlowBase

	Title       string
	Value       string
	Description string
}

type WorkFlowStepPermission struct {
	gorm.Model
	WorkFlowBase

	WorkFlowStepId int // FK
	CustomerId     int // FK
}

type WorkFlowPermission struct {
	gorm.Model
	WorkFlowBase

	WorkFlowId int // FK
	CustomerId int // FK
}
