package routes

import (
	"github.com/gorilla/mux"
	"book_api/pkg/controllers"
)

var RigisterBookStoreRoutes= func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}",controllers.GetByID).Methods("GET")
	router.HandleFunc("/book/{bookid}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/boo/{bookid}",controllers.DeleteBook).Methods("DELETE")
}