package datasource

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDbInstance struct {
	db           *gorm.DB
	dbName       string
	config       *gorm.Config
	initialDelay time.Duration
	debugLog     bool
}

func NewSqliteDataSource(sqliteDbFileName string) IDataSource {
	db := sqliteDbInstance{}
	db.config = &gorm.Config{}
	db.dbName = sqliteDbFileName
	db.initialDelay = 3 * time.Second

	db.connection(db.config)
	time.Sleep(db.initialDelay)
	db.ping()
	return &db
}

func (d *sqliteDbInstance) connection(config *gorm.Config) error {
	// d.dbName = d.dbName + "?cache=shared"
	d.dbName = d.dbName + "?busy_timeout=5000&journal_mode=WAL"
	db, err := gorm.Open(sqlite.Open(d.dbName), config)
	if err != nil {
		panic("failed to connect database")
	}
	d.db = db
	return nil
}

func (d *sqliteDbInstance) Disconnection() (disconnectError error) {
	fmt.Println("Disconnecting from the database…")
	sql, _ := d.db.DB()
	disconnectError = sql.Close()
	if disconnectError != nil {
		fmt.Println(disconnectError)
	}
	fmt.Println("Disconnected!")
	return nil
}

func (d *sqliteDbInstance) Migration(models []interface{}) error {
	d.db.AutoMigrate(models...)
	return nil
}

func (d *sqliteDbInstance) ping() (pingError error) {
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
