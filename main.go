package main

import (
	"emp-crud-api/controller"
	employeeentity "emp-crud-api/entity"
)

func main() {

	employeeentity.Employees = append(employeeentity.Employees, employeeentity.Employee{
		Id:        "1",
		EmpId:     1011,
		EmpName:   "Munish Kumar",
		CreatedBy: "admin",
		UpdatedBy: "editor",
	})

	controller.InitializeRouter()

}
