package payment

import (
	"github.com/mimani68/fintech-core/pkg/log"
)

type payment struct {
	log log.Ilogger
}

func PaymentHandler(logger log.Ilogger) payment {
	return payment{
		log: logger,
	}
}
