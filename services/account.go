package accountService

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	Debit  = "debit"
	Credit = "credit"
)

type transaction struct {
	ID            uuid.UUID
	EffectiveDate time.Time
	Type          string
	Description   string
	Amount        int
}

type transactionReduced struct {
	ID     uuid.UUID
	Type   string
	Amount int
}

type allTransactions []transaction
type allTransactionsReduced []transactionReduced

var actualBalance = 500
var transactions = allTransactions{
	{
		ID:            uuid.New(),
		Type:          Credit,
		Description:   "Credited salary",
		Amount:        500,
		EffectiveDate: time.Now(),
	},
}

func CreateTransaction(transactionType string, amount int, description string) (int, transaction, string) {
	var newTransaction transaction
	if (transactionType != Debit && transactionType != Credit) || amount <= 0 {
		return http.StatusBadRequest, transaction{}, "Invalid type of operation"
	} else if transactionType == Debit && actualBalance-amount < 0 {
		return http.StatusForbidden, transaction{}, "You don't have that amount of money available"
	} else {
		if transactionType == Credit {
			actualBalance += amount
		} else {
			actualBalance -= amount
		}
		newTransaction.Description = description
		newTransaction.Amount = amount
		newTransaction.Type = transactionType
		newTransaction.ID = uuid.New()
		newTransaction.EffectiveDate = time.Now()
		transactions = append(transactions, newTransaction)
		return http.StatusCreated, newTransaction, ""
	}
}

func GetTransactionDetails(transactionID uuid.UUID) (int, interface{}) {
	status := http.StatusNotFound
	var res interface{} = "Not found"
	for _, transactionLoop := range transactions {
		if transactionLoop.ID == transactionID {
			status = http.StatusOK
			res = transactionLoop
		}
	}
	return status, res
}

func GetBalance() (int, int) {
	return http.StatusOK, actualBalance
}

func GetAllTransactions() (int, interface{}) {
	reducedTransactionsResult := make(allTransactionsReduced, len(transactions))
	for i, item := range transactions {
		reducedTransactionsResult[i] = transactionReduced{
			ID:     item.ID,
			Amount: item.Amount,
			Type:   item.Type,
		}
	}
	return http.StatusOK, reducedTransactionsResult
}
