package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `josn:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main(){
	r:=mux.NewRouter()

	movies = append(movies,Movie{ID:"1",Isbn:"24779",Title:"Movie 1",Director:&Director{Firstname:"John",Lastname:"Doe"}})
	movies = append(movies,Movie{ID:"2",Isbn:"58600",Title:"Movie 2",Director:&Director{Firstname:"Trevor",Lastname:"Smith"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
    log.Fatal(http.ListenAndServe(":8000",r))
}

func getMovies(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)

	for _,item:=range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
	}
}
}

func deleteMovie(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Content-Type","application/json")

	params:=mux.Vars(req)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[index:],movies[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter ,req *http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie Movie

	_=json.NewDecoder(req.Body).Decode(&movie)

	movie.ID=strconv.Itoa(len(movies)+1)

	movies=append(movies,movie)
}

func updateMovie(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Content-Type","application/json")

	params:=mux.Vars(req)

	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[index:],movies[index+1:]...)

			var movie Movie

			_=json.NewDecoder(req.Body).Decode(&movie)

			movie.ID=params["id"]

			movies=append(movies,movie)

			json.NewEncoder(w).Encode(movie)
		}
	}
}