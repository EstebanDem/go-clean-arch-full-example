package usecases

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-clean-arch-example/internal/domain"
	"testing"
)

func TestAddEmployeeUseCaseInvalidRequest(t *testing.T) {
	uc := InitEmployeeUseCase(employeeRepoMock{})
	request := AddEmployeeRequest{
		Name:    "Michael",
		Country: "USA",
		Salary: SalaryRequest{
			Currency: "USD",
			Value:    0,
		},
	}
	_, err := uc.AddEmployee(context.Background(), request)

	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidValue, err)
}

func TestAddEmployeeUseCaseErrorOnSavingEmployeeRepo(t *testing.T) {
	employeeRepo := employeeRepoMock{}
	uc := InitEmployeeUseCase(&employeeRepo)
	request := AddEmployeeRequest{
		Name:    "Michael",
		Country: "USA",
		Salary: SalaryRequest{
			Currency: "USD",
			Value:    216,
		},
	}

	employeeRepo.On("Save", mock.Anything, mock.Anything).Return(errors.New("error saving employee"))
	_, err := uc.AddEmployee(context.Background(), request)

	assert.Error(t, err)
	assert.Equal(t, "error saving employee", err.Error())
}

func TestAddEmployeeUseCase(t *testing.T) {
	employeeRepo := employeeRepoMock{}
	uc := InitEmployeeUseCase(&employeeRepo)
	request := AddEmployeeRequest{
		Name:    "Michael",
		Country: "USA",
		Salary: SalaryRequest{
			Currency: "USD",
			Value:    216,
		},
	}

	employeeRepo.On("Save", mock.Anything, mock.Anything).Return(nil)
	_, err := uc.AddEmployee(context.Background(), request)

	assert.NoError(t, err)
}

// Mocks
type employeeRepoMock struct {
	mock.Mock
}

func (e2 employeeRepoMock) Save(ctx context.Context, e domain.Employee) error {
	args := e2.Called(ctx, e)
	return args.Error(0)
}

func (e2 employeeRepoMock) Delete(ctx context.Context, id uuid.UUID) error {
	args := e2.Called(ctx, id)
	return args.Error(0)
}

func (e2 employeeRepoMock) GetById(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	args := e2.Called(ctx, id)
	return args.Get(0).(*domain.Employee), args.Error(1)
}
