package repository

import (
	"context"
	"fmt"
	"muxMon/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type EmpRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmpRepo) InsertEmp(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *EmpRepo) FindEmpByID(empID string) (*model.Employee, error) {
	var emp model.Employee
	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Name: "employee_id", Value: empID}}).Decode(&emp)

	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmpRepo) FindAllEmp() ([]model.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []model.Employee

	if err = results.All(context.Background(), &emps); err != nil {
		return nil, fmt.Errorf("results decode error %s", err)
	}
	return emps, nil
}

func (r *EmpRepo) UpdateEmpID(empID string, updateEmp *model.Employee) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Name: "employee_id", Value: empID}},
		bson.D{{Name: "$set", Value: updateEmp}})

	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (r *EmpRepo) DeleteEmpByID(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Name: "employee_ID", Value: empID}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (r *EmpRepo) DeleteAllEmp() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
