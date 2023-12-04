package employeeentity

type Employee struct {
	Id        string `json:"id"`
	EmpId     int    `json:"empid"`
	EmpName   string `json:"empname"`
	CreatedBy string `json:"createdby"`
	UpdatedBy string `json:"updatedby"`
}

var Employees []Employee
