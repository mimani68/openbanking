package datasource

import "gorm.io/gorm"

type IDataSource interface {
	connection(config *gorm.Config) error
	Disconnection() error
	Migration(models []interface{}) error
	ping() (pingError error)
}
