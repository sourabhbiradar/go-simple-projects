package main

import (
	"fmt"
	"log"
	"html"
	"net/http"
)

func main(){
	http.HandleFunc("/",func (w http.ResponseWriter , r *http.Request){
		fmt.Fprintf(w,"Hello ,%q",html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/hi",func(w http.ResponseWriter ,r *http.Request){
		fmt.Fprintf(w,"Hi")
	})
	fmt.Println("Host port :8080 , Docker port :8081")
	log.Fatal(http.ListenAndServe(":8081",nil))
	
	
}