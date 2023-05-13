package payment

import (
	"github.com/mimani68/fintech-core/pkg/log"
	"github.com/mimani68/fintech-core/pkg/uow"
)

// example:
//
// TID := "jf93n22"
// trx := PaymentUnitOfWorkGenerator(TID, p.log)
//
//	ops := func() bool {
//		return false
//	}
//
//	rollbackOps := func() bool {
//		return true
//	}
//
// trx.Add(ops, rollbackOps)
// trx.Commit()
func PaymentUnitOfWorkGenerator(tid string, log log.Ilogger) PaymentUnitOfWork {
	log.Debug("Transaction Start", map[string]string{
		"TID": tid,
	})
	return PaymentUnitOfWork{
		Log: log,
		TID: tid,
		UnitOfWorkAbstract: uow.UnitOfWorkAbstract{
			TID: tid,
			Log: log,
		},
	}
}

type PaymentUnitOfWork struct {
	uow.UnitOfWorkAbstract

	TID       string
	isSuccess bool

	Log log.Ilogger
}

func (p *PaymentUnitOfWork) Add(cb func(map[string]interface{}) bool, rollback func(map[string]interface{}) bool, params map[string]interface{}) {
	p.isSuccess = cb(params)
	if p.isSuccess {
		p.Log.Debug("Add new ops in transition", map[string]string{
			"TID": p.TID,
		})
	} else {
		p.Log.Error("Rollback should applied", map[string]string{
			"TID": p.TID,
		})
		p.Rollback(rollback, params)
	}
}
