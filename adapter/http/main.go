package http

import (
	"net/http"

	"github.com/Viniciusmbc/dio-expert-session-finance/adapter/http/actuator"
	"github.com/Viniciusmbc/dio-expert-session-finance/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
