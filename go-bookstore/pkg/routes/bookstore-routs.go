package routes

import (
	"github.com/OgyreZ1/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegiterBookStoreRouts = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
}
