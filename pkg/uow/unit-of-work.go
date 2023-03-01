package uow

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
	// log                log.Ilogger
	transactionSuccess bool
}

func (u *UnitOfWorkAbstract) Commit() {
	// u.log.Debug("Commit", nil)
	u.transactionSuccess = true
}

func (u *UnitOfWorkAbstract) Rollback() {
	// u.log.Debug("Rollback", nil)
	u.transactionSuccess = false
}
