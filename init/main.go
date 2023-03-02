package main

import (
	"github.com/mimani68/fintech-core/core/payment"
	"github.com/mimani68/fintech-core/data/dto"
	. "github.com/mimani68/fintech-core/pkg/log"
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
	p := payment.PaymentHandler(generalLogger)
	p.PaymentDirect(&dto.PaymentRequestMeta{
		Amount:              10,
		DestinationBankCode: "15",
		IdempotencyId:       "ir2233",
	})
}
