package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/service/account"
	"github.com/cloudflare/cfssl/transport/roots/system"
	"github.com/labstack/echo"
	"github.com/mimani68/fintech-core/config"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
)

var serveCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  payment,
}

func serve(c *cli.Context) error {
	cfg := config.NewConfig()

	f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	mysqlRepo, err := mysql.New(cfg.Mysql, logger)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	accSrv := account.New(cfg.Account, mysqlRepo, logger)
	systemSrv := system.New(cfg.Account, mysqlRepo, logger)
	invitationSrv := invitation.New(cfg.Account, mysqlRepo, logger)

	restServer := echo.New(logger, accSrv, systemSrv, invitationSrv)
	go func() {
		if err := restServer.Start(cfg.App.Address); err != nil {
			logger.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	}

	return nil
}
