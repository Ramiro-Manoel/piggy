package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/finance_provider/pluggy"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/storage/postgres"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/handler"
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

	transactionRepo := postgres.NewAccountTransactionRepository(conn)
	categoryRepo := postgres.NewCategoryRepository(conn)
	accountRepo := postgres.NewAccountRepository(conn)

	financeProvider := pluggy.NewClient(os.Getenv("PLUGGY_CLIENT_ID"), os.Getenv("PLUGGY_CLIENT_SECRET"))
	err = financeProvider.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	transactionSvc := transaction.NewAccountService(transactionRepo, financeProvider)
	categorySvc := category.NewService(categoryRepo)
	accountSvc := account.NewService(accountRepo, financeProvider)

	handler := handler.NewHandler(transactionSvc, categorySvc, accountSvc, os.Getenv("PLUGGY_INSTITUTION_ID"))

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
