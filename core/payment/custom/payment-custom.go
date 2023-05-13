package custompayment

import (
	"fmt"

	"github.com/mimani68/fintech-core/core/payment"
	"github.com/mimani68/fintech-core/data/dto"
	"github.com/mimani68/fintech-core/pkg/flow"
	"github.com/mimani68/fintech-core/pkg/queue"
)

func (p *paymentAbstract) PaymentCustom(r *dto.PaymentRequestMeta) {
	p.Log.Debug("Incoming request for fulfillment was contain data like", map[string]string{
		"amount":      fmt.Sprintf("%v", r.Amount),
		"destination": fmt.Sprintf("%v", r.DestinationBankCode),
	})

	// Flow defining
	paymentFlow := flow.FlowGenerator(r.IdempotencyId, p.Log)

	// Validation
	paymentFlow.
		If("first-validation", firstValidation, "need-wait").
		Do("calculate-sth-1", func() {}).
		IfElse("second-validation", func() bool {
			return true
		}, "calculate-sth-2", "will-send-package")

	// Transaction Manager
	trx := payment.PaymentUnitOfWorkGenerator(r.IdempotencyId, p.Log)
	paymentFlow.Do("calculate-sth-2", func() {
		p.Log.Debug("Transaction manager", nil)
		// trx.Add(func() bool {
		// 	return true
		// }, func() bool {
		// 	return true
		// })
	}).
		If("simple-check", func() bool {
			return false
		}, "will-send-package")

	// Commit or Rollback
	paymentFlow.Do("calculate-sth-4", func() {
		trx.Commit()
	}).
		Do("need-wait", func() {
			p.Log.Debug("Finish flow by need-wait status", nil)
		}).
		Do("will-send-package", func() {
			p.Log.Debug("Finish flow by will-send-package", nil)
		})

	paymentFlow.Do("send-to-queue", func() {
		q := queue.QueueBuilder()
		paymentLabel := fmt.Sprintf("payment-%s", r.IdempotencyId)
		q.Add(paymentLabel, r)
	})

	paymentFlow.End()

}
