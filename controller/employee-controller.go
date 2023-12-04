package controller

import (
	employeeentity "emp-crud-api/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InitializeRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
	r.HandleFunc("/employees", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(employeeentity.Employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, employee := range employeeentity.Employees {
		if employee.Id == params["id"] {
			json.NewEncoder(w).Encode(employee)
			return
		}

	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"Error": "Employee not found"})

}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var employee employeeentity.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)

	id, _ := strconv.Atoi(employee.Id)

	employee.Id = fmt.Sprintf("%d", id+1)
	employeeentity.Employees = append(employeeentity.Employees, employee)
	json.NewEncoder(w).Encode(employee)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, employee := range employeeentity.Employees {
		if employee.Id == params["id"] {
			var updatedEmployee employeeentity.Employee
			json.NewDecoder(r.Body).Decode(&updatedEmployee)

			employee.EmpId = updatedEmployee.EmpId
			employee.EmpName = updatedEmployee.EmpName
			employee.UpdatedBy = updatedEmployee.UpdatedBy

			employeeentity.Employees[index] = employee

			json.NewEncoder(w).Encode(employee)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Employee not found"})

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, employee := range employeeentity.Employees {
		if employee.Id == params["id"] {
			employeeentity.Employees = append(employeeentity.Employees[:index], employeeentity.Employees[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Employee not found"})

}
