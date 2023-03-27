package view

import (
	"github.com/mimani68/fintech-core/data/model"
	"gorm.io/gorm"
)

func ViewCreator(db *gorm.DB) {
	cusotmerView := db.Model(&model.Customer{}).Select("users.id, users.name")
	db.Migrator().CreateView("v_customer", gorm.ViewOption{Query: cusotmerView})
}
