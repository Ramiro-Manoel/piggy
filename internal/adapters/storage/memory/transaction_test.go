package memory

import (
	"testing"
	"time"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func TestSaveAndRead(t *testing.T) {
	repo := NewTransactionRepository()
	transaction := transaction.AccountTransaction{}
	transaction.ID = "1"
	transaction.Description = "Mercado XYZ"
	transaction.Amount = 1050
	transaction.Date = time.Now()

	err := repo.Save(transaction)
	if err != nil {
		t.Fatalf("Save returned an error: %v", err)
	}

	transactionFound, err := repo.Read(transaction.ID)

	if err != nil {
		t.Fatalf("Read returned an error: %v", err)
	}
	if transactionFound != transaction {
		t.Errorf("expected %+v, recieved %+v", transaction, transactionFound)
	}
}

func TestReadNotFound(t *testing.T) {
	repo := NewTransactionRepository()

	_, err := repo.Read("non existing ID")

	if err == nil {
		t.Error("expected an error, did not recieved one")
	}
}
