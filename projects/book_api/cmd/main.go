package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	 "book_api/pkg/routes"
)

func main(){
	r:=mux.NewRouter()

	routes.RigisterBookStoreRoutes(r)

	http.Handle("/",r)

	log.Fatal(http.ListenAndServe(":9010",r))
}