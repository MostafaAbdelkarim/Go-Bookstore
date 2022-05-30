package routes

//importing controllers file to as it is the second stop after user sends requests through routes which then passed to model file then passed to DB through GORM
//gorilla/mux -> implements a request router and dispatcher for matching incoming requests to their respective handler
import (
	"github.com/MostafaAbdelkarim/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

//implementing a function that manages all routes of our backend API.
//every router.HandleFunc recieves a route and pass the request to a function implemented in controllers file with method type.
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
