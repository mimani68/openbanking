package model

import (
	"time"

	"gorm.io/gorm"
)

type PermissionBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Permission struct {
	gorm.Model
	PermissionBase

	Title string
	Value string
}

type RoleBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleAbstract struct {
	gorm.Model
	RoleBase

	Title          string
	Value          string
	PermissionList []Permission
}

type RoleIndevisual struct {
	gorm.Model
	RoleBase

	RoleId         int // FK
	IsCustomed     bool
	Title          string
	Value          string
	PermissionList []Permission
}

type RoleInstitutional struct {
	gorm.Model
	RoleBase

	RoleId         int // FK
	InstitueId     int // FK
	IsCustomed     bool
	Title          string
	Value          string
	PermissionList []Permission
}

// Polymorphic role handling
type RoleCustomer struct {
	gorm.Model
	RoleBase

	RoleType   string `default:"individual" json:"roleType" xml:"roleType" enum:"institutional,individual"`
	RoleId     int    // FK
	CustomerId int    // FK
}

type RoleCustomerActivityNotification struct {
	gorm.Model
	RoleBase

	Channel  string `default:"sms" json:"channel" xml:"channel" enum:"sms,slack,bale"`
	Values   []string
	Template string

	CriteriaType string `default:"sql" json:"criteriaType" xml:"criteriaType" enum:"sql,go"`
	Criteria     string // user=='admin'
}

type RoleCustomerHistory struct {
	gorm.Model
	RoleBase

	PerformerId int // FK
	Current     string
	Changed     string
}
