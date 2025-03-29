package employee

import (
	"context"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain/employee"
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
	AddEmployee(ctx context.Context, request AddEmployeeRequest) (AddEmployeeResponse, error)
}

type addEmployeeUseCase struct {
	employeesRepo employee.EmployeeRepository
}

func (uc addEmployeeUseCase) AddEmployee(ctx context.Context, request AddEmployeeRequest) (AddEmployeeResponse, error) {
	empl, err := employee.NewEmployee(request.Name, request.Country, request.Salary.Currency, request.Salary.Value)
	if err != nil {
		return AddEmployeeResponse{}, err
	}

	err = uc.employeesRepo.Save(ctx, *empl)
	if err != nil {
		return AddEmployeeResponse{}, err
	}

	return AddEmployeeResponse{
		Id:        empl.Id,
		CreatedAt: empl.CreatedAt,
	}, nil
}

func InitEmployeeUseCase(er employee.EmployeeRepository) AddEmployeeUseCase {
	return addEmployeeUseCase{
		employeesRepo: er,
	}
}
