package main

import (
	"fmt"

	"github.com/mimani68/fintech-core/core/payment/direct"
	"github.com/mimani68/fintech-core/data/dto"
	. "github.com/mimani68/fintech-core/pkg/log"
	"github.com/mimani68/fintech-core/pkg/random"
)

func main() {

	// Logger instace generator
	generalLogger := Log()
	generalLogger.Info("App is about to start", nil)
	generalLogger.Info("Running", map[string]string{
		"developer-mode": "True",
	})
	generalLogger.Debug("Current version for steady connectin in all distribution is LTS 11", map[string]string{
		"version": "1",
	})
	generalLogger.Error("Port definition", map[string]string{
		"expected": "3000",
		"current":  "NULL",
	})

	// Run simple payment operation
	p := direct.PaymentHandler(generalLogger)
	p.PaymentDirect(&dto.PaymentRequestMeta{
		Amount:              10e3 * 2,
		DestinationBankCode: "shr100-337-10084-1",
		IdempotencyId:       fmt.Sprintf("ir%d", random.RandomInt(10e3, 10e6)),
	})
}
