package payments

// Payment e interface za razlicni plakanja
type Payment interface {
	ProcessPayment() string
}