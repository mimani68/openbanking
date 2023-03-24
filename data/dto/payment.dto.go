package dto

import "github.com/mimani68/fintech-core/pkg/tracer"

type PaymentRequest struct {
	Amount                 int
	DestinationBankCode    string
	DestinationBankId      int
	DestinationAccountCode string
	DestinationAccountId   int
}

type PaymentRequestMeta struct {
	PaymentRequest

	CallbackWebHook string

	IdempotencyId string
	CustomerId    string
	ServerId      string

	Tracer tracer.ITracer
}

type PaymentResponseDto struct {
	Message string `json:"message"`
	TID     string `json:"tid"`
}
