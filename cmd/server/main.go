package main

import (
	"log"
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/storage/memory"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/handler"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func main() {

	transactionRepo := memory.NewTransactionRepository()
	categoryRepo := memory.NewCategoryRepository()
	accountRepo := memory.NewAccountRepository()

	transactionSvc := transaction.NewService(transactionRepo)
	categorySvc := category.NewService(categoryRepo)
	accountSvc := account.NewService(accountRepo)

	handler := handler.NewHandler(transactionSvc, categorySvc, accountSvc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
