package main

import (
	"log"
	"net/http"

	"github.com/OgyreZ1/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := mux.NewRouter()
	routes.RegiterBookStoreRouts(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
