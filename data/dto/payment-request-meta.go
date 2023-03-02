package dto

import "github.com/mimani68/fintech-core/pkg/tracer"

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
