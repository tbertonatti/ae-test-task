package accountController

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	service "github.com/tbertonatti/ae-test-task/services"
)

type transactionDTO struct {
	Type        string `json:"Type"`
	Description string `json:"Description"`
	Amount      int    `json:"Amount"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction transactionDTO
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the transaction information")
	}
	json.Unmarshal(reqBody, &newTransaction)
	status, res, errorR := service.CreateTransaction(newTransaction.Type, newTransaction.Amount, newTransaction.Description)
	w.WriteHeader(status)
	if len(errorR) == 0 {
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(errorR)
	}
}

func GetTransactionDetails(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
		}
	}()
	transactionID := uuid.MustParse(mux.Vars(r)["id"])
	state, res := service.GetTransactionDetails(transactionID)
	w.WriteHeader(state)
	json.NewEncoder(w).Encode(res)
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	state, res := service.GetBalance()
	w.WriteHeader(state)
	json.NewEncoder(w).Encode(res)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	state, res := service.GetAllTransactions()
	w.WriteHeader(state)
	json.NewEncoder(w).Encode(res)
}
