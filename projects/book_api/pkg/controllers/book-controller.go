package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"book_api/pkg/utils"
	"book_api/pkg/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter , r *http.Request){
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetByID(w http.ResponseWriter , r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]

	ID,err:=strconv.ParseInt(bookId,0,0)

	if err!=nil{
		fmt.Println("Error while Parsing")
	}

	bookDetails,_:=models.GetByID(ID)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter , r *http.Request){
	createBook:=&models.Book{}
	utils.ParseBody(r,createBook)

	b:=createBook.CreateBook()
	res,_:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId:=vars["bookId"]

	ID,err:=strconv.ParseInt(bookId,0,0)

	if err!=nil{
		fmt.Println("Error While Parsing")

	}
	book:=models.DeleteBook(ID)

	res,_:=json.Marshal(book)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook =&models.Book{}
	utils.ParseBody(r,updateBook)

	vars :=mux.Vars(r)
	bookId:=vars["bookId"]

	ID,err:=strconv.ParseInt(bookId,0,0)

	if err!=nil{
		fmt.Println("Error while Parsing")

	}
	bookDetails,db:=models.GetByID(ID)
	if updateBook.Name!=""{
		bookDetails.Name=updateBook.Name
	}
	if updateBook.Author!=""{
		bookDetails.Author=updateBook.Author
	}
	if updateBook.Publications!=""{
		bookDetails.Publications=updateBook.Publications
	}

	db.Save(&bookDetails)

	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}