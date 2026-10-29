package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/finance_provider/pluggy"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/storage/postgres"
	"github.com/Ramiro-Manoel/piggy/internal/card"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/handler"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	accountTransactionRepo := postgres.NewAccountTransactionRepository(pool)
	cardRepo := postgres.NewCardRepository(pool)
	invoiceRepo := postgres.NewInvoiceRepository(pool)
	cardTransactionRepo := postgres.NewCardTransactionRepository(pool)
	categoryRepo := postgres.NewCategoryRepository(pool)
	accountRepo := postgres.NewAccountRepository(pool)

	financeProvider := pluggy.NewClient(os.Getenv("PLUGGY_CLIENT_ID"), os.Getenv("PLUGGY_CLIENT_SECRET"))
	err = financeProvider.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	accountSvc := account.NewService(accountRepo, financeProvider)
	cardSvc := card.NewService(cardRepo, financeProvider)
	invoiceSvc := invoice.NewService(invoiceRepo)
	accountTransactionSvc := transaction.NewAccountService(accountTransactionRepo, financeProvider, accountSvc)
	cardTransactionSvc := transaction.NewCardService(cardTransactionRepo, financeProvider, cardSvc, invoiceSvc)
	categorySvc := category.NewService(categoryRepo)

	handler := handler.NewHandler(accountTransactionSvc, cardTransactionSvc, categorySvc, accountSvc, cardSvc, os.Getenv("PLUGGY_INSTITUTION_ID"))

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
