package usecase

import (
	"encoding/json"
	"log"
	"muxMon/model"
	"muxMon/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmpService struct {
	MongoCollection *mongo.Collection
}

type Respond struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func (srv *EmpService) CreateEmp(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Respond{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee
	err := json.NewDecoder(req.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}
	emp.EmpID = uuid.NewString()
	repo := repository.EmpRepo{MongoCollection: srv.MongoCollection}

	insertID, err := repo.InsertEmp(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp.EmpID
	w.WriteHeader(http.StatusOK)
	log.Println("Emp created with ID", insertID, emp)
}

func (srv *EmpService) GetEmpByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Respond{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(req)["id"]
	log.Println("employee id", empID)

	repo := repository.EmpRepo{MongoCollection: srv.MongoCollection}
	emp, err := repo.FindEmpByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error :", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (srv *EmpService) GetAllEmps(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Respond{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmpRepo{MongoCollection: srv.MongoCollection}
	emps, err := repo.FindAllEmp()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error :", err)
		res.Error = err.Error()
		return
	}
	res.Data = emps
	w.WriteHeader(http.StatusOK)
}

func (srv *EmpService) UpdateEmpByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Respond{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(req)["id"]
	log.Println("employee id", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid id")
		res.Error = "invalid id"
		return
	}
	var emp model.Employee

	err := json.NewDecoder(req.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}
	emp.EmpID = empID
	repo := repository.EmpRepo{MongoCollection: srv.MongoCollection}

	count, err := repo.UpdateEmpID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error :", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (srv *EmpService) DeleteEmpByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Respond{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(req)["id"]
	log.Println("employee id", empID)

	repo := repository.EmpRepo{MongoCollection: srv.MongoCollection}
	count, err := repo.DeleteEmpByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error :", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (srv *EmpService) DeleteAllEmps(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Respond{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmpRepo{MongoCollection: srv.MongoCollection}
	count, err := repo.DeleteAllEmp()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error :", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
