package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)

	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form",formHandler)

	fmt.Println("starting server at port :8080\n")

	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter,r *http.Request){
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParseForm() error :%v",err)
		return
	}
	fmt.Fprintf(w,"POST request successful\n")
	name:=r.FormValue("name")
	address:=r.FormValue("address")

	fmt.Fprintf(w , "Name : %s\n",name)
	fmt.Fprintf(w,"Address : %s\n",address)
}

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"Method is not support",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"Hello!")
}