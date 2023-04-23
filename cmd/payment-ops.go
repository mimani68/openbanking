package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mimani68/fintech-core/core/payment/direct"
	"github.com/mimani68/fintech-core/data/datasource"
	"github.com/mimani68/fintech-core/data/dto"
	"github.com/mimani68/fintech-core/pkg/log"
	"github.com/mimani68/fintech-core/pkg/random"
	"github.com/urfave/cli/v2"
)

var paymentCMD = &cli.Command{
	Name:    "payment",
	Aliases: []string{"p"},
	Usage:   "Performing payments operations on your data",
	Action:  development,
}

func payment(c *cli.Context) error {
	// cfg := config.NewConfig()

	// Logger instance generator
	generalLogger := log.Log()
	generalLogger.Info("App is about to start", nil)
	generalLogger.Info("Running", map[string]string{
		"developer-mode": "True",
	})
	generalLogger.Debug("Current version for steady connection in all distribution is LTS 11", map[string]string{
		"version": "1",
	})
	generalLogger.Error("Port definition", map[string]string{
		"expected": "3000",
		"current":  "NULL",
	})

	db := datasource.NewSqliteDataSource("sample.db")
	fmt.Printf("%T\n", db)

	// Run simple payment operation
	p := direct.PaymentHandler(generalLogger)
	p.PaymentDirect(&dto.PaymentRequestMeta{
		PaymentRequest: dto.PaymentRequest{
			Amount:              10e3 * 2,
			DestinationBankCode: "shr100-337-10084-1",
		},
		IdempotencyId: fmt.Sprintf("ir%d", random.RandomInt(10e3, 10e6)),
	})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	return nil
}
