package main

import (
	"fmt"
	"net/http"
	"github.com/louahola/cdis/repository"
	"log"
	"encoding/json"
	"github.com/louahola/cdis/factory"
)

var repo repository.Repository

func handlerICon(w http.ResponseWriter, r *http.Request) {}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	repo = new(repository.MongoRepository)

	http.HandleFunc("/favicon", handlerICon)

	http.HandleFunc("/", handler)
	http.HandleFunc("/index", handler)

	http.HandleFunc("/generateMonthlyReport", generateMonthlyReportHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}