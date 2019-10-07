package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing Request....!")
	w.Write([]byte("OK"))
	log.Println("Finished Procesing Request")
}

func main() {
	fmt.Println("Server is running in port 8000...")

	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggedRouter)
}
