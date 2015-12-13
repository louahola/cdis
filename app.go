package main

import (
	"net/http"
	"github.com/louahola/cdis/repository"
	"github.com/louahola/cdis/api"
	"log"
	"github.com/louahola/cdis/web"
)


func main() {
	var repo repository.Repository = new(repository.MongoRepository)

	api.Initialize(repo)
	web.Initialize(repo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}