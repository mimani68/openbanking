package model

import (
	"time"

	"gorm.io/gorm"
)

type WorkflowBase struct {
	Id        int       `json:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"index,not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"index"`
}

type WorkflowAbstract struct {
	gorm.Model
	WorkflowBase

	Title       string `json:"title" gorm:"index,not null"`
	Description string `json:"description"`

	BPMNFile   string `json:"bpmnFile"`
	HasDiagram bool   `json:"hasDiagram" default:"false"`
}

type Workflow struct {
	gorm.Model
	WorkflowBase

	WorkflowAbstract   []WorkflowAbstract `json:"workflowAbstract" xml:"workflowAbstract" gorm:"foreignKey:WorkflowAbstractId"`
	WorkflowAbstractId int                `json:"workflowAbstractId" xml:"workflowAbstractId"`
	Payment            Payment            `json:"payment" xml:"payment" gorm:"foreignKey:PaymentId"`
	PaymentId          int                `json:"paymentId" xml:"paymentId" `

	Status WorkflowStatus
}

type WorkflowStatus struct {
	gorm.Model
	WorkflowBase

	Workflow   []Workflow `json:"workflow" xml:"workflow" gorm:"foreignKey:workflowId"`
	WorkflowId int        `json:"workflowId" xml:"workflowId"`
	Status     string     `default:"PENDING" json:"status" xml:"status" enum:"DONE,PENDING"`
}

type WorkflowStep struct {
	gorm.Model
	WorkflowBase

	Title       string `json:"title" xml:"title" gorm:"index"`
	Value       string `json:"value" xml:"value"`
	Order       int    `json:"order" xml:"order"`
	Description string `json:"description" xml:"description"`
}

type WorkflowPermissionBase struct {
	Id        int       `json:"id" gorm:"primaryKey,unique,not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"index,not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"index"`
}

type WorkflowStepPermissionType struct {
	gorm.Model
	WorkflowPermissionBase

	Title string `json:"title" xml:"title" gorm:"index,not null"`
	Value string `json:"value" xml:"value" gorm:"index,not null"`
}

type WorkflowStepPermission struct {
	gorm.Model
	WorkflowPermissionBase

	WorkflowStep   []WorkflowStep               `json:"workflowStep" xml:"workflowStep" gorm:"foreignKey:workflowStepId"`
	WorkflowStepId int                          `json:"workflowStepId" xml:"workflowStepId"`
	Customer       []Customer                   `json:"customer" xml:"customer" gorm:"foreignKey:customerId"`
	CustomerId     int                          `json:"customerId" xml:"customerId"`
	Permission     []WorkflowStepPermissionType `json:"permission" xml:"permission" gorm:"foreignKey:permissionId"`
	PermissionId   int                          `json:"permissionId" xml:"permissionId"`
}

type WorkflowPermissionType struct {
	gorm.Model
	WorkflowPermissionBase

	Title string `json:"title" xml:"title" gorm:"index,not null"`
	Value string `json:"value" xml:"value" gorm:"index,not null"`
}

type WorkflowPermission struct {
	gorm.Model
	WorkflowPermissionBase

	Workflow     []Workflow               `json:"workflow" xml:"workflow" gorm:"foreignKey:workflowId"`
	WorkflowId   int                      `json:"workflowId" xml:"workflowId"`
	Customer     []Customer               `json:"customer" xml:"customer" gorm:"foreignKey:customerId"`
	CustomerId   int                      `json:"customerId" xml:"customerId"`
	Permission   []WorkflowPermissionType `json:"permission" xml:"permission" gorm:"foreignKey:permissionId"`
	PermissionId int                      `json:"permissionId" xml:"permissionId"`
}
