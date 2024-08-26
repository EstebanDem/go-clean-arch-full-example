package usecases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-clean-arch-example/internal/domain"
	"go-clean-arch-example/internal/pkg"
	"testing"
)

func TestGetEmployeeSalaryUseCase_GetSalaryEmployeeNotFound(t *testing.T) {
	employeeRepo := employeeRepoMock{}
	request := GetEmployeeSalaryRequest{
		EmployeeId: pkg.GenerateUUID(),
		Currency:   "MXN",
	}

	employeeRepo.On("GetById", mock.Anything, mock.Anything).Return(&domain.Employee{}, domain.ErrEmployeeNotFound)
	uc := InitGetEmployeeSalaryUseCase(&employeeRepo, &currencyConverterMock{})

	_, err := uc.GetSalary(context.Background(), request)
	assert.Equal(t, domain.ErrEmployeeNotFound, err)
}

func TestGetEmployeeSalaryUseCase_GetSalaryEmployeeRatioServiceError(t *testing.T) {
	employeeRepo := employeeRepoMock{}
	currencyConverter := currencyConverterMock{}
	request := GetEmployeeSalaryRequest{
		EmployeeId: pkg.GenerateUUID(),
		Currency:   "MXN",
	}

	employeeRepo.On("GetById", mock.Anything, mock.Anything).Return(&domain.Employee{
		Salary: domain.Salary{
			Currency: "USD",
			Value:    3000,
		},
	}, nil)
	currencyConverter.On("GetExchangeRate", mock.Anything, mock.Anything).Return(0.0, errors.New("error in service"))

	uc := InitGetEmployeeSalaryUseCase(&employeeRepo, &currencyConverter)

	_, err := uc.GetSalary(context.Background(), request)
	assert.Equal(t, "error in service", err.Error())
}

func TestGetEmployeeSalaryUseCase_GetSalaryEmployee(t *testing.T) {
	employeeRepo := employeeRepoMock{}
	currencyConverter := currencyConverterMock{}
	request := GetEmployeeSalaryRequest{
		EmployeeId: pkg.GenerateUUID(),
		Currency:   "MXN",
	}

	employeeRepo.On("GetById", mock.Anything, mock.Anything).Return(&domain.Employee{
		Salary: domain.Salary{
			Currency: "USD",
			Value:    3000,
		},
	}, nil)
	currencyConverter.On("GetExchangeRate", mock.Anything, mock.Anything).Return(0.051, nil)

	uc := InitGetEmployeeSalaryUseCase(&employeeRepo, &currencyConverter)

	resp, err := uc.GetSalary(context.Background(), request)

	assert.NoError(t, err)
	assert.Equal(t, 0.051, resp.ConvertedSalary.Rate)
	assert.Equal(t, 3000.0, resp.Salary.Value)
	assert.Equal(t, 153.0, resp.ConvertedSalary.Value)
}

type currencyConverterMock struct {
	mock.Mock
}

func (m *currencyConverterMock) GetExchangeRate(baseCurrency string, currency string) (float64, error) {
	args := m.Called(baseCurrency, currency)
	return args.Get(0).(float64), args.Error(1)
}
