package uow

import (
	"fmt"

	"github.com/mimani68/fintech-core/pkg/log"
)

/**
  # Unit of Work
  A unit of work is a pattern used to group
  one or more operations into a single unit
  so that all the operations either pass or fail.
  Say user registration, all the transaction insert/update/delete
  and so on done in one single transaction
  rather than done in multiple database transactions.
*/
type UnitOfWorkAbstract struct {
	Log log.Ilogger

	TID                string
	transactionSuccess bool
}

func (u *UnitOfWorkAbstract) Commit() {
	u.transactionSuccess = true
	u.Log.Debug("Commit", map[string]string{
		"TID":    u.TID,
		"status": fmt.Sprintf("%v", u.transactionSuccess),
	})
}

func (u *UnitOfWorkAbstract) Rollback(cb func(map[string]interface{}) bool, params map[string]interface{}) {
	state := cb(params)
	if state {
		u.transactionSuccess = false
		u.Log.Debug("Rollback", map[string]string{
			"TID":    u.TID,
			"status": fmt.Sprintf("%v", u.transactionSuccess),
		})
	} else {
		u.Log.Error("Rollback operation was unsuccessful", map[string]string{
			"TID":    u.TID,
			"status": fmt.Sprintf("%v", u.transactionSuccess),
		})
	}
}
