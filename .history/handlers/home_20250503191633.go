package handlers

import (
	"go-payment-app-basik/payments"
	"html/template"
	"net/http"
)

// HomeHandler прикажува Home Page и обработува избрано плаќање
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Ако методот на барањето е POST (корисникот испратил форма)
	if r.Method == http.MethodPost {
		r.ParseForm() // Го парсира формуларот од HTTP барањето (ги чита сите вредности од формата)
		paymentType := r.FormValue("payment") // Ја зема избраната вредност од полето "payment" од формата

		var p payments.Payment // Декларира интерфејс од тип Payment од пакетот payments

		// Врз основа на избраниот тип на плаќање, иницијализира соодветна структура
		switch paymentType {
		case "card":
			p = payments.CardPayment{} // Ако е избрано "card", се користи CardPayment структурата
		case "cash":
			p = payments.CashPayment{} // Ако е избрано "cash", се користи CashPayment структурата
		case "paypal":
			p = payments.PayPalPayment{} // Ако е избрано "paypal", се користи PayPalPayment структурата
		default:
			http.Error(w, "Invalid payment method", http.StatusBadRequest) // Ако е нешто друго, враќа HTTP грешка 400
			return // Прекинува понатамошно извршување
		}

		result := p.ProcessPayment() // Го повикува методот ProcessPayment од соодветната структура и го чува резултатот

		tmpl := template.Must(template.ParseFiles("templates/index.html")) // Ја вчитува HTML шаблоната од датотеката index.html
		tmpl.Execute(w, map[string]string{"Result": result}) // Ја прикажува шаблоната и го испраќа резултатот до клиентот
		return // Прекинува понатамошно извршување
	}

	// Ако е GET барање - само ја прикажува формата без обработка
	tmpl := template.Must(template.ParseFiles("templates/index.html")) // Ја вчитува шаблоната
	tmpl.Execute(w, nil) // Ја прикажува на клиентот без дополнителни податоци
}
