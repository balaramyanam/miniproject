package main

import (
	"log"

	"net/http"

	"github.com/balaram/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	coder := mux.NewRouter()
	routes.RsgisterBookstoreRoutes(coder)
	http.Handle("/", coder)
	
	log.Fatal(http.ListenAndServe("localhost:7070", coder))
}
