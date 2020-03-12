package models

//to read from redis
//E are capital because json marshall has to access them
type Employee struct {
	EmployeeID string `json:"employeeID"`
	EmployeeName string `json:"employeeName"`
}

var Employees []Employee


type NewEmployeeResponse struct {
	Source string `json:"source"`
	EmployeeName string `json:"employeeName"`
}