package dto

type PayeeDto struct {
	PersonDto int

	DestinationBankCode    string
	DestinationBankId      int
	DestinationAccountCode string
	DestinationAccountId   int
}
