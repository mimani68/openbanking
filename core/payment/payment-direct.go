package payment

import (
	"fmt"

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
}
