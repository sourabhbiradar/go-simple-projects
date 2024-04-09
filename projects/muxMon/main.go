package main

import (
	"context"
	"log"
	"muxMon/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// load .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("could not load .env")
	}
	log.Println(".env Loaded")

	// create mongo client
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("connection failed", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("mongo connected")
}

func main() {
	// close db connection
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empService := usecase.EmpService{MongoCollection: coll}

	// api layer
	r := mux.NewRouter()

	r.HandleFunc("/employee", empService.CreateEmp).Methods("POST")
	r.HandleFunc("/employee/{id}", empService.GetEmpByID).Methods(http.MethodGet)
	r.HandleFunc("/employee", empService.GetAllEmps).Methods("GET")
	r.HandleFunc("/employee/{id}", empService.UpdateEmpByID).Methods("PUT")
	r.HandleFunc("/employee/{id}", empService.DeleteEmpByID).Methods("DELETE")
	r.HandleFunc("/employee", empService.DeleteAllEmps).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
