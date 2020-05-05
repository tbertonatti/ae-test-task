package main

import (
	"log"
	"net/http"

	routes "github.com/tbertonatti/ae-test-task/routes"
)

func main() {
	// Starting backend api
	router := routes.AccountRoutes()
	http.Handle("/api/", router)
	// Starting frontend
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	log.Printf("App starting in http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
