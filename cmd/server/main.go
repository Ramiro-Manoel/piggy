package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/finance_provider/pluggy"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/storage/memory"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/storage/postgres"
	"github.com/Ramiro-Manoel/piggy/internal/card"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/handler"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	accountTransactionRepo := postgres.NewAccountTransactionRepository(conn)
	cardRepo := postgres.NewCardRepository(conn)
	invoiceRepo := memory.NewInvoiceRepository()
	cardTransactionRepo := postgres.NewCardTransactionRepository(conn)
	categoryRepo := postgres.NewCategoryRepository(conn)
	accountRepo := postgres.NewAccountRepository(conn)

	financeProvider := pluggy.NewClient(os.Getenv("PLUGGY_CLIENT_ID"), os.Getenv("PLUGGY_CLIENT_SECRET"))
	err = financeProvider.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	accountSvc := account.NewService(accountRepo, financeProvider)
	cardSvc := card.NewService(cardRepo, financeProvider)
	invoiceSvc := invoice.NewService(invoiceRepo)
	transactionSvc := transaction.NewAccountService(accountTransactionRepo, financeProvider)
	cardTransactionSvc := transaction.NewCardService(cardTransactionRepo, financeProvider, cardSvc, invoiceSvc)
	categorySvc := category.NewService(categoryRepo)

	handler := handler.NewHandler(transactionSvc, cardTransactionSvc, categorySvc, accountSvc, cardSvc, os.Getenv("PLUGGY_INSTITUTION_ID"))

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
