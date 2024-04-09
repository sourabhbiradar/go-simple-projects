package model

type Employee struct {
	EmpID      string `json:"emp_id,omitempty" bson:"emp_id"`
	Name       string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}
