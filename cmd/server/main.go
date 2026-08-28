package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/adapters/storage/memory"
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

	transactionRepo := memory.NewTransactionRepository()
	categoryRepo := memory.NewCategoryRepository()
	accountRepo := postgres.NewAccountRepository(conn)

	transactionSvc := transaction.NewService(transactionRepo)
	categorySvc := category.NewService(categoryRepo)
	accountSvc := account.NewService(accountRepo)

	handler := handler.NewHandler(transactionSvc, categorySvc, accountSvc)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
