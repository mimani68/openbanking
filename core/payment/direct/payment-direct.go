package directpayment

import (
	"fmt"

	"github.com/mimani68/fintech-core/core/payment"
	"github.com/mimani68/fintech-core/data/dto"
	"github.com/mimani68/fintech-core/pkg/flow"
	"github.com/mimani68/fintech-core/pkg/queue"
	"github.com/mimani68/fintech-core/policy"
)

//
// Adding current payment request to operation queue
//
func (p *paymentAbstract) PaymentDirect(r *dto.PaymentRequestMeta) {
	p.Log.Debug("Incoming request for fullfilment was contain data like", map[string]string{
		"amount":      fmt.Sprintf("%v", r.Amount),
		"destination": fmt.Sprintf("%v", r.DestinationBankCode),
	})

	// Flow defining
	paymentFlow := flow.FlowGenerator(r.IdempotencyId, p.Log)

	// Validation
	paymentFlow.
		IfElse("minimum-payment-value-meet", func() bool {
			p.Log.Debug("Minimum payment limit must meet", r)
			return r.Amount >= policy.PaymentPolicy.MinumumPayment
		}, "maxium-payment-value-meet", "send-to-queue").
		IfElse("maxium-payment-value-meet", func() bool {
			p.Log.Debug("Maximum payment limit must meet", r)
			return r.Amount < policy.PaymentPolicy.MaxiumPayment
		}, "payer-balance-adequacy", "send-to-queue").
		IfElse("payer-balance-adequacy", func() bool {
			payerBalance := 100
			p.Log.Debug("Payer bank account balance adequacy must meet", r)
			return r.Amount >= payerBalance
		}, "payer-operation-state", "send-to-queue").
		IfElse("payer-operation-state", func() bool {
			payerState := true
			p.Log.Debug("Payer should allowed operation", r)
			return payerState
		}, "do-payer-allowed-to-send-to-payee", "send-to-queue").
		IfElse("do-payer-allowed-to-send-to-payee", func() bool {
			payerState := true
			p.Log.Debug("Payer must allowed for transfer", r)
			return payerState
		}, "payment-transaction-start", "send-to-queue")

	// Transaction Start
	var trx payment.PaymentUnitOfWork
	paymentFlow.Do("payment-transaction-start", func() {
		trx = payment.PaymentUnitOfWorkGenerator(r.IdempotencyId, p.Log)
		p.Log.Debug("Transaction manager started", nil)
	})

	// Attachment of payment
	paymentFlow.Do("attachment", func() {
		p.Log.Debug("Attach if needed", nil)
		fileName := "doc-1"
		trx.Add(attachFileIntoStorage(fileName), removeFileFromStorage(), fileName)
	})

	paymentFlow.Do("send-to-queue", func() {
		q := queue.QueueBuilder()
		paymentLabel := fmt.Sprintf("payment-%s", r.IdempotencyId)
		q.Add(paymentLabel, r)
	})

	paymentFlow.End()

}
