package direct

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
func (p *paymentAbstract) PaymentDirect(r *dto.PaymentRequestMeta) dto.PaymentResonseDto {
	p.Log.Debug("Incoming request for fullfilment was contain data like", map[string]string{
		"amount":      fmt.Sprintf("%v", r.Amount),
		"destination": fmt.Sprintf("%v", r.DestinationBankCode),
	})
	var response dto.PaymentResonseDto

	// Flow defining
	paymentFlow := flow.FlowGenerator(r.IdempotencyId, p.Log)

	// Validation
	paymentFlow.
		IfElse("minimum-payment-value-meet", func() bool {
			p.Log.Debug("Minimum payment limit must meet", map[string]string{
				"amount":          fmt.Sprintf("%v", r.Amount),
				"minimum-allowed": fmt.Sprintf("%v", policy.PaymentPolicy.MinumumPayment),
			})
			return r.Amount > policy.PaymentPolicy.MinumumPayment
		}, "maxium-payment-value-meet", "reject-payment").
		IfElse("maxium-payment-value-meet", func() bool {
			p.Log.Debug("Maximum payment limit must meet", map[string]string{
				"amount":          fmt.Sprintf("%v", r.Amount),
				"maximum-allowed": fmt.Sprintf("%v", policy.PaymentPolicy.MaxiumPayment),
			})
			return r.Amount < policy.PaymentPolicy.MaxiumPayment
		}, "payer-balance-adequacy", "reject-payment").
		IfElse("payer-balance-adequacy", func() bool {
			payerBalance := 100
			p.Log.Debug("Payer bank account balance adequacy must meet", map[string]string{
				"amount":        fmt.Sprintf("%v", r.Amount),
				"payer-balance": fmt.Sprintf("%v", payerBalance),
			})
			return r.Amount >= payerBalance
		}, "payer-operation-state", "reject-payment")

	// Policy checker
	paymentFlow.
		IfElse("payer-operation-state", func() bool {
			payerState := true
			p.Log.Debug("Payer should allowed operation", nil)
			return payerState
		}, "do-payer-allowed-to-send-to-payee", "reject-payment").
		IfElse("do-payer-allowed-to-send-to-payee", func() bool {
			payerState := true
			p.Log.Debug("Payer must allowed for transfer", nil)
			return payerState
		}, "is-payment-method-activated", "reject-payment")

	// Payment policy checker
	paymentFlow.
		IfElse("is-payment-method-activated", func() bool {
			payerState := true
			p.Log.Debug("Payer pay using permited transfer method like SWIFT or SHABA", nil)
			return payerState
		}, "payment-transaction-start", "reject-payment")

	// Transaction Start
	var trx payment.PaymentUnitOfWork
	paymentFlow.Do("payment-transaction-start", func() {
		trx = payment.PaymentUnitOfWorkGenerator(r.IdempotencyId, p.Log)
		p.Log.Debug("Transaction manager started", nil)
	})

	// Attachment of payment
	paymentFlow.Do("attachment", func() {
		p.Log.Debug("Attach if needed", map[string]string{
			"filename": "doc-1",
			"TID":      r.IdempotencyId,
		})
		file := map[string]interface{}{
			"file": "doc-1",
		}
		trx.Add(attachFileIntoStorage, removeFileFromStorage, file)
	})

	// Reject payment request
	paymentFlow.Do("reject-payment", func() {
		response = dto.PaymentResonseDto{
			Message: "Due failure in pre requirements the transaction had failed",
			TID:     r.IdempotencyId,
		}
	}, "attachment")

	// Send data into event queue
	paymentFlow.Do("send-to-queue", func() {
		q := queue.QueueBuilder()
		paymentLabel := fmt.Sprintf("payment-%s", r.IdempotencyId)
		q.Add(paymentLabel, r)
		response = dto.PaymentResonseDto{
			Message: "Transaction queued successfuly and will wait for further process",
			TID:     r.IdempotencyId,
		}
	}, "reject-payment")

	paymentFlow.End()

	p.Log.Debug("Final response", map[string]string{
		"message": response.Message,
		"TID":     response.TID,
	})
	return response

}
