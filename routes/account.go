package accountRoute

import (
	"github.com/gorilla/mux"
	controller "github.com/tbertonatti/ae-test-task/controllers"
)

func AccountRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/balance", controller.GetBalance).Methods("GET")
	router.HandleFunc("/api/transactions", controller.CreateTransaction).Methods("POST")
	router.HandleFunc("/api/transactions", controller.GetAllTransactions).Methods("GET")
	router.HandleFunc("/api/transactions/{id}", controller.GetTransactionDetails).Methods("GET")
	return router
}
