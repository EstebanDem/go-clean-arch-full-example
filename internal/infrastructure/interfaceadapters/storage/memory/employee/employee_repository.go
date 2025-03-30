package employee

import (
	"context"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain/employee"
	"time"
)

type InMemoryEmployee struct {
	Id        uuid.UUID
	Name      string
	Country   string
	Salary    InMemorySalary
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InMemorySalary struct {
	Currency string
	Value    float64
}

func toDomainEmployee(emp InMemoryEmployee) employee.Employee {
	return employee.Employee{
		Id:      emp.Id,
		Name:    emp.Name,
		Country: emp.Country,
		Salary: employee.Salary{
			Currency: emp.Salary.Currency,
			Value:    emp.Salary.Value,
		},
		CreatedAt: emp.CreatedAt,
		UpdatedAt: emp.CreatedAt,
	}
}

func toInMemoryEmployee(e employee.Employee) InMemoryEmployee {
	return InMemoryEmployee{
		Id:      e.Id,
		Name:    e.Name,
		Country: e.Country,
		Salary: InMemorySalary{
			Currency: e.Salary.Currency,
			Value:    e.Salary.Value,
		},
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

type InMemoryEmployeeRepository struct {
	employees map[uuid.UUID]InMemoryEmployee
}

func NewInMemoryEmployeeRepository() InMemoryEmployeeRepository {
	employees := make(map[uuid.UUID]InMemoryEmployee)

	// default employee added for testing
	eId, _ := uuid.Parse("aa02193c-0592-4191-955f-eefdc04ea35d")
	now := time.Now()
	employees[eId] = InMemoryEmployee{
		Id:      eId,
		Name:    "George",
		Country: "Michael",
		Salary: InMemorySalary{
			Currency: "USD",
			Value:    1000,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	return InMemoryEmployeeRepository{
		employees: employees,
	}
}

func (i InMemoryEmployeeRepository) Save(_ context.Context, e employee.Employee) error {
	employee := toInMemoryEmployee(e)
	i.employees[employee.Id] = employee
	return nil
}

func (i InMemoryEmployeeRepository) Delete(_ context.Context, id uuid.UUID) error {
	_, ok := i.employees[id]
	if !ok {
		return employee.ErrEmployeeNotFound
	}

	delete(i.employees, id)
	return nil
}

func (i InMemoryEmployeeRepository) GetById(_ context.Context, id uuid.UUID) (*employee.Employee, error) {
	empl, ok := i.employees[id]
	if !ok {
		return nil, employee.ErrEmployeeNotFound
	}

	dEmployee := toDomainEmployee(empl)

	return &dEmployee, nil
}
