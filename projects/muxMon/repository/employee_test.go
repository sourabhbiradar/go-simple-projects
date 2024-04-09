package repository

import (
	"context"
	"log"
	"muxMon/model"
	"testing"

	"github.com/google/uuid"

	//"go.mongodb.org/mongo-driver/internal/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {

	connectStr := "mongodb+srv://muxCRUD:muxCRUD8@cluster8.rwostqv.mongodb.net/?retryWrites=true&w=majority&appName=Cluster8"

	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectStr))

	if err != nil {
		log.Fatal("error while connecting mongoDB", err)
	}

	log.Println("mongoDB connected")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("Ping successful")
	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dummy data
	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	// connect to collection
	coll := mongoTestClient.Database("companydb").Collection("employee_test")
	empRepo := EmpRepo{MongoCollection: coll}

	// insert emp 1 data

	t.Run("Insert emp 1", func(t *testing.T) {
		emp := model.Employee{Name: "Abc", Department: "Dev", EmpID: emp1}
		result, err := empRepo.InsertEmp(&emp)
		if err != nil {
			t.Fatal("Insert 1 failed", err)
		}
		t.Log("Insert 1 successful", result)
	})

	// insert emp 2
	t.Run("Insert emp 2", func(t *testing.T) {
		emp := model.Employee{Name: "Xyz", Department: "HR", EmpID: emp2}
		result, err := empRepo.InsertEmp(&emp)
		if err != nil {
			t.Fatal("Insert 2 failed", err)
		}
		t.Log("Insert 2 successful", result)
	})

	// get emp 1 data
	t.Run("Get emp 1 data", func(t *testing.T) {
		result, err := empRepo.FindEmpByID(emp1)
		if err != nil {
			t.Fatal("failed to get emp 1", err)
		}
		t.Log("emp1", result.Name)
	})

	// get all emps
	t.Run("Get all emps data", func(t *testing.T) {
		results, err := empRepo.FindAllEmp()

		if err != nil {
			t.Fatal("failed to get all emps", err)
		}
		t.Log("Emps", results)
	})

	//  update emp 1 data
	t.Run("Update emp 1 data", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Abd",
			Department: "Sr Dev",
			EmpID:      emp1,
		}
		result, err := empRepo.UpdateEmpID(emp1, &emp)
		if err != nil {
			t.Fatal("update failed", err)
		}
		t.Log("update successfull", result)

	})

	// get emp 1 after update
	t.Run("Get emp 1 after update", func(t *testing.T) {
		result, err := empRepo.FindEmpByID(emp1)
		if err != nil {
			t.Fatal("failed to get emp 1", err)
		}
		t.Log("emp1", result.Name)
	})

	// delete emp 1
	t.Run("Delete emp 1", func(t *testing.T) {
		result, err := empRepo.DeleteEmpByID(emp1)
		if err != nil {
			t.Fatal("failed to delete", err)
		}
		t.Log("delete count", result)
	})

	// get all emps
	t.Run("Get all emps data", func(t *testing.T) {
		results, err := empRepo.FindAllEmp()

		if err != nil {
			t.Fatal("failed to get all emps", err)
		}
		t.Log("Emps", results)
	})

	// delete all emps
	t.Run("Delete all emps", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmp()
		if err != nil {
			t.Fatal("failed to delete all emps", err)
		}
		t.Log("delete count", result)
	})

}
