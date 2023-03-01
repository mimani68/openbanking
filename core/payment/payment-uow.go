package payment

import (
	"github.com/mimani68/fintech-core/pkg/log"
	"github.com/mimani68/fintech-core/pkg/uow"
)

func PaymentUnitOfWorkGenerator(log log.Ilogger) PaymentUnitOfWork {
	log.Debug("Transaction Start", nil)
	return PaymentUnitOfWork{
		log: log,
	}
}

type PaymentUnitOfWork struct {
	uow.UnitOfWorkAbstract
	log log.Ilogger
}

func (p *PaymentUnitOfWork) Add(cb func()) {
	p.log.Debug("Add new ops in transtion", nil)
}
