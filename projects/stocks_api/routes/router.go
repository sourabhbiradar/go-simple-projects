package routes

import (
	"github.com/gorilla/mux"
	"stocks_api/middleware"
)

func Router() *mux.Router{
	router:=mux.NewRouter()

	router.HandleFunc("/api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stocks",middleware.GetAllStocks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newstock",middleware.CreateStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deletestock/{id}",middleware.DeleteStock).Methods("DELETE","OPTIONS")
    return router
}