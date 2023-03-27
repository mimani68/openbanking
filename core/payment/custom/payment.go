package custompayment

import (
	"github.com/mimani68/fintech-core/pkg/log"
)

type paymentAbstract struct {
	Log log.Ilogger
}

func CustomPaymentHandler(logger log.Ilogger) paymentAbstract {
	return paymentAbstract{
		Log: logger,
	}
}
