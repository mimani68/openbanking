package policy

type PaymentPolicyAbstract struct {
	MinumumPayment int
	MaxiumPayment  int
}

var PaymentPolicy = &PaymentPolicyAbstract{
	MinumumPayment: 10e3,
	MaxiumPayment:  10e7,
}
