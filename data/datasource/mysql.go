package datasource

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDbInstance struct {
	db           *gorm.DB
	dbName       string
	config       *gorm.Config
	initialDelay time.Duration
	debugLog     bool
}

func NewMySqlDataSource(dbAddress string) IDataSource {
	db := mysqlDbInstance{}
	db.config = &gorm.Config{}
	db.dbName = dbAddress
	db.initialDelay = 3 * time.Second
	db.debugLog = false

	db.connection(db.config)
	time.Sleep(db.initialDelay)
	db.ping()
	return &db
}

func (d *mysqlDbInstance) connection(config *gorm.Config) error {
	if d.debugLog {
		d.dbName = d.dbName + "?charset=utf8&timezone=Asia/Tehran"
	}
	db, err := gorm.Open(mysql.Open(d.dbName), config)
	if err != nil {
		panic("failed to connect database")
	}
	d.db = db
	return nil
}

func (d *mysqlDbInstance) Disconnection() (disconnectError error) {
	fmt.Println("Disconnecting from the database…")
	sql, _ := d.db.DB()
	disconnectError = sql.Close()
	if disconnectError != nil {
		fmt.Println(disconnectError)
	}
	fmt.Println("Disconnected!")
	return nil
}

func (d *mysqlDbInstance) Migration(models []interface{}) error {
	d.db.AutoMigrate(models...)
	return nil
}

func (d *mysqlDbInstance) ping() (pingError error) {
	sql, err := d.db.DB()
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		return err
	} else {
		fmt.Println("Database initial ping")
	}

	pingError = sql.Ping()
	if pingError != nil {
		fmt.Println("Error pinging database: ", pingError)
		return pingError
	} else {
		fmt.Println("Database connection is still alive")
		return nil
	}

}
