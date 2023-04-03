package datasource

import "gorm.io/gorm"

type IDataSource interface {
	connection(config *gorm.Config) error
	disconnection() error
	migrations(models interface{}) error
	ping() (pingError error)
}
