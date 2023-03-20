package directpayment

import (
	"github.com/mimani68/fintech-core/pkg/log"
)

type paymentAbstract struct {
	Log log.Ilogger
}

func PaymentHandler(logger log.Ilogger) paymentAbstract {
	return paymentAbstract{
		Log: logger,
	}
}
