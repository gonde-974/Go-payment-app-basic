package payments

//CashPayment e struktura za pplakanje vo kes

type CashPayment struct {}

func (c CashPayment) ProcessPayment() string{
	return "Uspesno plakanje so kes"
}