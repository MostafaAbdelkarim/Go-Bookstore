package main

import (
	"log"
	"net/http"

	"github.com/MostafaAbdelkarim/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Main needs a main function
func main() {
	//Linking the main file to the Routes
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	//logging errors if happened, passing ListenAndServe Function to run the server on port 9010
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
