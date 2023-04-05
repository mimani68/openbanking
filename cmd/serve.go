package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
)

var serveCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  payment,
}

func serve(c *cli.Context) error {
	// cfg := config.NewConfig()

	// db := datasource.NewSqliteDataSource("sample.db")

	// restServer := echo.NewEchoApp(db, cfg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	// if err := restServer.Shutdown(); err != nil {
	// 	fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	// }

	return nil
}
