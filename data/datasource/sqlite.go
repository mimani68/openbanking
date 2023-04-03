package datasource

import (
	"fmt"
	"time"

	"github.com/mimani68/fintech-core/data/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type databaseInstance struct {
	db           *gorm.DB
	dbName       string
	config       *gorm.Config
	initialDelay time.Duration
}

func NewSqliteDataSource(sqliteDbFileName string) IDataSource {
	db := databaseInstance{}
	db.config = &gorm.Config{}
	db.dbName = sqliteDbFileName
	db.initialDelay = 3 * time.Second

	db.connection(db.config)
	time.Sleep(db.initialDelay)
	db.ping()
	db.migrations(
		model.Account{},
	)
	return &db
}

func (d *databaseInstance) connection(config *gorm.Config) error {
	db, err := gorm.Open(sqlite.Open(d.dbName), config)
	if err != nil {
		panic("failed to connect database")
	}
	d.db = db
	return nil
}

func (d *databaseInstance) disconnection() (disconnectError error) {
	fmt.Println("Disconnecting from the database…")
	sql, _ := d.db.DB()
	disconnectError = sql.Close()
	if disconnectError != nil {
		fmt.Println(disconnectError)
	}
	fmt.Println("Disconnected!")
	return nil
}

func (d *databaseInstance) migrations(models interface{}) error {
	d.db.AutoMigrate(models)
	return nil
}

func (d *databaseInstance) ping() (pingError error) {
	sql, err := d.db.DB()
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		return err
	} else {
		fmt.Println("Database connection is still alive")
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
