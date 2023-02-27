package payment

import (
	. "github.com/mimani68/fintech-core/pkg/log"
)

type payment struct {
	log Ilogger
}

func PaymentHandler(logger Ilogger) payment {
	return payment{
		log: logger,
	}
}
