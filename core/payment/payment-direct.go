package payment

import "fmt"

type PaymentRequestMeta struct {
	Amount                 int
	DestinationBankCode    string
	DestinationBankId      int
	DestinationAccountCode string
	DestinationAccountId   int
}

func (p *payment) PaymentDirect(r *PaymentRequestMeta) {
	p.log.Debug("Incoming request for fullfilment was contain data like", map[string]string{
		"amount":      fmt.Sprintf("%v", r.Amount),
		"destination": fmt.Sprintf("%v", r.DestinationBankCode),
	})
}
