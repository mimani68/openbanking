package policy

type PaymentPolicyAbstract struct {
	MinimumPayment int
	MaximumPayment int
}

var PaymentPolicy = &PaymentPolicyAbstract{
	MinimumPayment: 10e3,
	MaximumPayment: 10e7,
}
