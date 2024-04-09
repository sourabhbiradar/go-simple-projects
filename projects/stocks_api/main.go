package main

import (
	"net/http"
	"fmt"
	"stocks_api/routes"
	"log"
	_ "github.com/lib/pq"
)

func main(){
	r:=routes.Router()

	fmt.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080",r))
}