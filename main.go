package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting...")
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", HOME)
	// http.ListenAndServe(":8080", router)
	http.ListenAndServe(":"+port, router)
}

func HOME(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em Execução!")
}
