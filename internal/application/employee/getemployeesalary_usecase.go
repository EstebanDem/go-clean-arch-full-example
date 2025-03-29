package employee

import (
	"context"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/application/currency"
	"go-clean-arch-example/internal/domain/employee"
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
	employeesRepo     employee.EmployeeRepository
	currencyConverter currency.CurrencyConverter
}

func (g getEmployeeSalaryUseCase) GetSalary(ctx context.Context, request GetEmployeeSalaryRequest) (GetEmployeeSalaryResponse, error) {
	empl, err := g.employeesRepo.GetById(ctx, request.EmployeeId)
	if err != nil {
		return GetEmployeeSalaryResponse{}, err
	}

	ratio, err := g.currencyConverter.GetExchangeRate(empl.Salary.Currency, request.Currency)
	if err != nil {
		return GetEmployeeSalaryResponse{}, err
	}

	convertedSalary := ratio * empl.Salary.Value

	return GetEmployeeSalaryResponse{
		EmployeeId: empl.Id,
		Salary: Salary{
			Currency: empl.Salary.Currency,
			Value:    empl.Salary.Value,
		},
		ConvertedSalary: ConvertedSalary{
			Currency: request.Currency,
			Value:    convertedSalary,
			Rate:     ratio,
		},
	}, nil

}

func InitGetEmployeeSalaryUseCase(er employee.EmployeeRepository, cc currency.CurrencyConverter) GetEmployeeSalaryUseCase {
	return getEmployeeSalaryUseCase{
		employeesRepo:     er,
		currencyConverter: cc,
	}
}
