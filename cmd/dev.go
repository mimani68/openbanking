package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mimani68/fintech-core/data/datasource"
	"github.com/mimani68/fintech-core/data/model"
	"github.com/urfave/cli/v2"
)

var devCMD = &cli.Command{
	Name:    "dev",
	Aliases: []string{"d"},
	Usage:   "development process",
	Action:  development,
}

func development(c *cli.Context) error {
	// cfg := config.NewConfig()

	// Logger instance generator
	// generalLogger := log.Log()

	databaseAddress := "root:123@tcp(127.0.0.1:3306)/app"
	db := datasource.NewMySqlDataSource(databaseAddress)

	db.Migration([]interface{}{
		&model.Person{},
		&model.PersonMeta{},
	})
	// time.Sleep(2 * time.Second)
	// db.Disconnection()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	return nil
}
