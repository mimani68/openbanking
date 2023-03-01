package payment

import (
	"fmt"

	"github.com/mimani68/fintech-core/pkg/flow"
	"github.com/mimani68/fintech-core/pkg/tracer"
)

type PaymentRequestMeta struct {
	Amount                 int
	DestinationBankCode    string
	DestinationBankId      int
	DestinationAccountCode string
	DestinationAccountId   int

	IdempotencyId string
	CustomerId    string
	ServerId      string

	Tracer tracer.ITracer
}

func (p *payment) PaymentDirect(r *PaymentRequestMeta) {
	p.log.Debug("Incoming request for fullfilment was contain data like", map[string]string{
		"amount":      fmt.Sprintf("%v", r.Amount),
		"destination": fmt.Sprintf("%v", r.DestinationBankCode),
	})

	// Flow defining
	paymentFlow := flow.FlowGenerator(r.IdempotencyId, p.log)
	// Validation
	paymentFlow.
		If("first-validation", firstValidation, "need-wait").
		Do("calculate-sth-1", func() {}).
		IfElse("second-validation", func() bool {
			return true
		}, "calculate-sth-2", "will-send-package")
	// Transaction Manager
	trx := PaymentUnitOfWorkGenerator(p.log)
	paymentFlow.Do("calculate-sth-2", func() {
		p.log.Debug("Transaction manager", nil)
		trx.Add(func() {})
		trx.Add(func() {})
	}).
		If("simple-check", func() bool {
			return false
		}, "will-send-package").
		Do("calculate-sth-3", func() {
			trx.Add(func() {})
			trx.Add(func() {})
		})
	// Commit or Rollback
	paymentFlow.Do("calculate-sth-4", func() {
		trx.Commit()
	}).
		Do("need-wait", func() {
			p.log.Debug("Finish flow by need-wait status", nil)
		}).
		Do("will-send-package", func() {
			p.log.Debug("Finish flow by will-send-package", nil)
		})

}

func firstValidation() bool {
	return true
}
