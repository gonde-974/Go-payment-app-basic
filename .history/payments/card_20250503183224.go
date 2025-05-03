package payments

// Card e struktura za plakanje so karticka

type CardPayment struct {}

func (c CardPayment) ProcessPayment()string{
	return "Uspesno plateno so kreditna karticka."
}