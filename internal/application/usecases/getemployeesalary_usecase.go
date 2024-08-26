package usecases

import (
	"context"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/application/services"
	"go-clean-arch-example/internal/domain"
)

type GetEmployeeSalaryRequest struct {
	EmployeeId uuid.UUID
	Currency   string
}

type GetEmployeeSalaryResponse struct {
	EmployeeId      uuid.UUID
	Salary          Salary
	ConvertedSalary ConvertedSalary
}

type Salary struct {
	Currency string
	Value    float64
}

type ConvertedSalary struct {
	Currency string
	Value    float64
	Rate     float64
}

type GetEmployeeSalaryUseCase interface {
	GetSalary(ctx context.Context, request GetEmployeeSalaryRequest) (GetEmployeeSalaryResponse, error)
}

type getEmployeeSalaryUseCase struct {
	employeesRepo     domain.EmployeeRepository
	currencyConverter services.CurrencyConverter
}

func (g getEmployeeSalaryUseCase) GetSalary(ctx context.Context, request GetEmployeeSalaryRequest) (GetEmployeeSalaryResponse, error) {
	employee, err := g.employeesRepo.GetById(ctx, request.EmployeeId)
	if err != nil {
		return GetEmployeeSalaryResponse{}, err
	}

	ratio, err := g.currencyConverter.GetExchangeRate(employee.Salary.Currency, request.Currency)
	if err != nil {
		return GetEmployeeSalaryResponse{}, err
	}

	convertedSalary := ratio * employee.Salary.Value

	return GetEmployeeSalaryResponse{
		EmployeeId: employee.Id,
		Salary: Salary{
			Currency: employee.Salary.Currency,
			Value:    employee.Salary.Value,
		},
		ConvertedSalary: ConvertedSalary{
			Currency: request.Currency,
			Value:    convertedSalary,
			Rate:     ratio,
		},
	}, nil

}

func InitGetEmployeeSalaryUseCase(er domain.EmployeeRepository, cc services.CurrencyConverter) GetEmployeeSalaryUseCase {
	return getEmployeeSalaryUseCase{
		employeesRepo:     er,
		currencyConverter: cc,
	}
}
