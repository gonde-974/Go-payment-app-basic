package handlers

import (
	"go-payment-app-basik/payments"
	"html/template"
	"net/http"
)

// HomeHandler прикажува Home Page и обработува избрано плаќање
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		paymentType := r.FormValue("payment")

		var p payments.Payment

		switch paymentType {
		case "card":
			p = payments.CardPayment{}
		case "cash":
			p = payments.CashPayment{}
		case "paypal":
			p = payments.PayPalPayment{}
		default:
			http.Error(w, "Invalid payment method", http.StatusBadRequest)
			return
		}

		result := p.ProcessPayment()
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, map[string]string{"Result": result})
		return
	}

	// GET метод - само прикажи формата
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}
