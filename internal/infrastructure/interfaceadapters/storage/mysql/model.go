package mysql

import (
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
	"time"
)

// EmployeeRecord represents each row in 'employee' table
type EmployeeRecord struct {
	id        uint64
	uuid      uuid.UUID
	Name      string
	SalaryId  uint64 //FK
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SalaryRecord represents each row in 'salary' table
type SalaryRecord struct {
	id       uint64
	Currency string
	Wage     float64
}

// EmployeeWithSalary represents the result after joining tables
type EmployeeWithSalary struct {
	id        uint64 // is it needed?
	uuid      uuid.UUID
	Name      string
	Country   string
	Currency  string
	Wage      float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toEmployeeDomain(er EmployeeWithSalary) domain.Employee {
	return domain.Employee{
		Id:      er.uuid,
		Name:    er.Name,
		Country: er.Country,
		Salary: domain.Salary{
			Currency: er.Currency,
			Value:    er.Wage,
		},
		CreatedAt: er.CreatedAt,
		UpdatedAt: er.UpdatedAt,
	}
}
