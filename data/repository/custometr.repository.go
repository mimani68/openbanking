package repository

import (
	"errors"
	"time"

	"github.com/mimani68/fintech-core/data/model"
	"gorm.io/gorm"
)

type CustomerServiceRepo struct {
	db *gorm.DB
}

func NewCustomerServiceRepo(db *gorm.DB) *CustomerServiceRepo {
	return &CustomerServiceRepo{
		db: db,
	}
}

func (r *CustomerServiceRepo) Create(customerService *model.CustomerService) error {
	result := r.db.Create(customerService)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerServiceRepo) GetSingleItem(serviceId int) (*model.CustomerService, error) {
	var customerService model.CustomerService
	result := r.db.Preload("Service").First(&customerService, "service_id = ?", serviceId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer service not found")
		}
		return nil, result.Error
	}
	return &customerService, nil
}

func (r *CustomerServiceRepo) GetListOfItems(limit, offset int) ([]*model.CustomerService, error) {
	var customerServices []*model.CustomerService
	result := r.db.Preload("Service").Limit(limit).Offset(offset).Find(&customerServices)
	if result.Error != nil {
		return nil, result.Error
	}
	return customerServices, nil
}

func (r *CustomerServiceRepo) GetByQuery(query string) ([]*model.CustomerService, error) {
	var customerServices []*model.CustomerService
	result := r.db.Preload("Service").Where(query).Find(&customerServices)
	if result.Error != nil {
		return nil, result.Error
	}
	return customerServices, nil
}

func (r *CustomerServiceRepo) UpdateSingleAttribute(serviceId int, attribute string, value interface{}) error {
	result := r.db.Model(&model.CustomerService{}).Where("service_id = ?", serviceId).Update(attribute, value)
	if result.RowsAffected == 0 {
		return errors.New("customer service not found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerServiceRepo) UpdateWholeItem(customerService *model.CustomerService) error {
	result := r.db.Save(customerService)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerServiceRepo) SoftDelete(serviceId int) error {
	result := r.db.Model(&model.CustomerService{}).Where("service_id = ?", serviceId).Update("deleted_at", time.Now())
	if result.RowsAffected == 0 {
		return errors.New("customer service not found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerServiceRepo) HardDelete(serviceId int) error {
	result := r.db.Where("service_id = ?", serviceId).Delete(&model.CustomerService{})
	if result.RowsAffected == 0 {
		return errors.New("customer service not found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerServiceRepo) GroupCreate(customerServices []*model.CustomerService) error {
	result := r.db.CreateInBatches(customerServices, len(customerServices))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerServiceRepo) GroupDelete(serviceIds []int) error {
	result := r.db.Where("service_id IN (?)", serviceIds).Delete(&model.CustomerService{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
