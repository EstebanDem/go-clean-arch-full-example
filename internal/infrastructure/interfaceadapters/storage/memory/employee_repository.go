package memory

import (
	"context"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
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

func toDomainEmployee(emp InMemoryEmployee) domain.Employee {
	return domain.Employee{
		Id:      emp.Id,
		Name:    emp.Name,
		Country: emp.Country,
		Salary: domain.Salary{
			Currency: emp.Salary.Currency,
			Value:    emp.Salary.Value,
		},
		CreatedAt: emp.CreatedAt,
		UpdatedAt: emp.CreatedAt,
	}
}

func toInMemoryEmployee(e domain.Employee) InMemoryEmployee {
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
	return InMemoryEmployeeRepository{
		employees: employees,
	}
}

func (i InMemoryEmployeeRepository) Save(_ context.Context, e domain.Employee) error {
	employee := toInMemoryEmployee(e)
	i.employees[employee.Id] = employee
	return nil
}

func (i InMemoryEmployeeRepository) Delete(_ context.Context, id uuid.UUID) error {
	_, ok := i.employees[id]
	if !ok {
		return domain.ErrEmployeeNotFound
	}

	delete(i.employees, id)
	return nil
}

func (i InMemoryEmployeeRepository) GetById(_ context.Context, id uuid.UUID) (*domain.Employee, error) {
	employee, ok := i.employees[id]
	if !ok {
		return nil, domain.ErrEmployeeNotFound
	}

	dEmployee := toDomainEmployee(employee)

	return &dEmployee, nil
}
