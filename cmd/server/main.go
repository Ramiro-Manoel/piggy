package main

import (
	"log"
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/adapters/storage/memory"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/handler"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func main() {

	transactionRepo := memory.NewTransactionRepository()
	categoryRepo := memory.NewCategoryRepository()

	transactionSvc := transaction.NewService(transactionRepo)
	categorySvc := category.NewService(categoryRepo)

	handler := handler.NewHandler(transactionSvc, categorySvc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
