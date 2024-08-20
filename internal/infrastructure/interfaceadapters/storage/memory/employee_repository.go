package memory

import (
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
	"time"
)

type InMemoryEmployee struct {
	Id        uuid.UUID
	Name      string
	Country   string
	SalaryId  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (i InMemoryEmployeeRepository) toDomainEmployee(emp InMemoryEmployee) (domain.Employee, error) {
	salary, err := i.salaryRepo.GetById(emp.SalaryId)
	if err != nil {
		return domain.Employee{}, err
	}

	return domain.Employee{
		Id:        emp.Id,
		Name:      emp.Name,
		Country:   emp.Country,
		Salary:    salary,
		CreatedAt: emp.CreatedAt,
		UpdatedAt: emp.CreatedAt,
	}, nil
}

func toInMemoryEmployee(e domain.Employee) InMemoryEmployee {
	return InMemoryEmployee{
		Id:        e.Id,
		Name:      e.Name,
		Country:   e.Country,
		SalaryId:  e.Salary.Id,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

type InMemoryEmployeeRepository struct {
	employees  map[uuid.UUID]InMemoryEmployee
	salaryRepo InMemorySalaryRepository
}

func NewInMemoryEmployeeRepository(salaryRepo InMemorySalaryRepository) InMemoryEmployeeRepository {
	employees := make(map[uuid.UUID]InMemoryEmployee)
	return InMemoryEmployeeRepository{
		employees:  employees,
		salaryRepo: salaryRepo,
	}
}

func (i InMemoryEmployeeRepository) Save(e domain.Employee) error {
	employee := toInMemoryEmployee(e)
	_, err := i.salaryRepo.GetById(e.Salary.Id)
	if err != nil {
		return err
	}
	i.employees[employee.Id] = employee
	return nil
}

func (i InMemoryEmployeeRepository) Delete(id uuid.UUID) error {
	_, ok := i.employees[id]
	if !ok {
		return domain.ErrEmployeeNotFound
	}

	delete(i.employees, id)
	return nil
}

func (i InMemoryEmployeeRepository) GetById(id uuid.UUID) (*domain.Employee, error) {
	employee, ok := i.employees[id]
	if !ok {
		return nil, domain.ErrEmployeeNotFound
	}

	dEmployee, err := i.toDomainEmployee(employee)
	if err != nil {
		return nil, err
	}

	return &dEmployee, nil
}
