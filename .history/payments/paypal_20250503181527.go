package payments

//PayPalPayment struktura za plakane so paypal

type PayPalPayment struct {}

func (p PayPalPayment) ProcessPayment() string{
	return "Uspesno plaknke so paypal"
}