package usecases

import (
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
	"time"
)

type AddEmployeeRequest struct {
	Name    string
	Country string
	Salary  SalaryRequest
}

type SalaryRequest struct {
	Currency string
	Value    float64
}

type AddEmployeeResponse struct {
	Id        uuid.UUID
	CreatedAt time.Time
}

type AddEmployeeUseCase interface {
	AddEmployee(request AddEmployeeRequest) (AddEmployeeResponse, error)
}

type addEmployeeUseCase struct {
	employeesRepo domain.EmployeeRepository
	salaryRepo    domain.SalaryRepository
}

func (uc addEmployeeUseCase) AddEmployee(request AddEmployeeRequest) (AddEmployeeResponse, error) {
	employee, err := domain.NewEmployee(request.Name, request.Country, request.Salary.Currency, request.Salary.Value)
	if err != nil {
		return AddEmployeeResponse{}, err
	}

	err = uc.employeesRepo.Save(*employee)
	if err != nil {
		return AddEmployeeResponse{}, err
	}

	err = uc.salaryRepo.Save(*employee.Salary)
	if err != nil {
		return AddEmployeeResponse{}, err
	}

	return AddEmployeeResponse{
		Id:        employee.Id,
		CreatedAt: employee.CreatedAt,
	}, nil
}

func InitEmployeeUseCase(er domain.EmployeeRepository, sr domain.SalaryRepository) AddEmployeeUseCase {
	return addEmployeeUseCase{
		employeesRepo: er,
		salaryRepo:    sr,
	}
}
