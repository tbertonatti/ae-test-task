package accountService

import (
	"log"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestGetAllTransactions(t *testing.T) {
	status, transactionsResult := GetAllTransactions()
	if status == http.StatusOK {
		log.Print(transactionsResult)
	} else {
		t.Errorf("GetAllTransactions() == %d, want %d", status, http.StatusOK)
	}
}

func TestGetBalance(t *testing.T) {
	status, balance := GetBalance()
	if status != http.StatusOK {
		t.Errorf("GetBalance() == %d, want %d", status, http.StatusOK)
	}
	if balance != 500 {
		t.Errorf("GetBalance() == %d, want %d", balance, 500)
	}
}

func TestCreateTransaction(t *testing.T) {
	status, newTransaction, errorR := CreateTransaction(Debit, 5, "prueba")
	if status != http.StatusCreated {
		t.Errorf(errorR)
		t.Errorf("CreateTransaction() == %d, want %d", status, http.StatusCreated)
	} else {
		status, transactionR := GetTransactionDetails(newTransaction.ID)
		if status != http.StatusOK {
			t.Errorf("GetTransactionDetails() == %d, want %d", status, http.StatusOK)
		} else {
			log.Print(transactionR)
		}
	}
}

func TestGet404TransRecord(t *testing.T) {
	var id uuid.UUID
	status, res := GetTransactionDetails(id)
	if status != http.StatusNotFound {
		t.Errorf("CreateTransaction() == %d, want %d", status, http.StatusNotFound)
	} else {
		log.Print(res)
	}
}
func TestGetNewBalance(t *testing.T) {
	status, balance := GetBalance()
	if status != http.StatusOK {
		t.Errorf("GetBalance() after debit of 5 == %d, want %d", status, http.StatusOK)
	}
	if balance != 495 {
		t.Errorf("GetBalance() after debit of 5 == %d, want %d", balance, 495)
	}
}
