package usecases

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-clean-arch-example/internal/domain"
	"testing"
)

func TestAddEmployeeUseCaseInvalidRequest(t *testing.T) {
	uc := InitEmployeeUseCase(employeeRepoMock{}, salaryRepoMock{})
	request := AddEmployeeRequest{
		Name:    "Michael",
		Country: "USA",
		Salary: SalaryRequest{
			Currency: "USD",
			Value:    0,
		},
	}
	_, err := uc.AddEmployee(request)

	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidValue, err)
}

func TestAddEmployeeUseCaseErrorOnSavingSalaryRepo(t *testing.T) {
	salaryRepo := salaryRepoMock{}
	uc := InitEmployeeUseCase(employeeRepoMock{}, &salaryRepo)
	request := AddEmployeeRequest{
		Name:    "Michael",
		Country: "USA",
		Salary: SalaryRequest{
			Currency: "USD",
			Value:    216,
		},
	}

	salaryRepo.On("Save", mock.Anything).Return(errors.New("error saving salary"))
	_, err := uc.AddEmployee(request)

	assert.Error(t, err)
	assert.Equal(t, "error saving salary", err.Error())
}

func TestAddEmployeeUseCaseErrorOnSavingEmployeeRepo(t *testing.T) {
	salaryRepo := salaryRepoMock{}
	employeeRepo := employeeRepoMock{}
	uc := InitEmployeeUseCase(&employeeRepo, &salaryRepo)
	request := AddEmployeeRequest{
		Name:    "Michael",
		Country: "USA",
		Salary: SalaryRequest{
			Currency: "USD",
			Value:    216,
		},
	}

	salaryRepo.On("Save", mock.Anything).Return(nil)
	employeeRepo.On("Save", mock.Anything).Return(errors.New("error saving employee"))
	_, err := uc.AddEmployee(request)

	assert.Error(t, err)
	assert.Equal(t, "error saving employee", err.Error())
}

func TestAddEmployeeUseCase(t *testing.T) {
	salaryRepo := salaryRepoMock{}
	employeeRepo := employeeRepoMock{}
	uc := InitEmployeeUseCase(&employeeRepo, &salaryRepo)
	request := AddEmployeeRequest{
		Name:    "Michael",
		Country: "USA",
		Salary: SalaryRequest{
			Currency: "USD",
			Value:    216,
		},
	}

	salaryRepo.On("Save", mock.Anything).Return(nil)
	employeeRepo.On("Save", mock.Anything).Return(nil)
	_, err := uc.AddEmployee(request)

	assert.NoError(t, err)
}

// Mocks
type employeeRepoMock struct {
	mock.Mock
}

func (e2 employeeRepoMock) Save(e domain.Employee) error {
	args := e2.Called(e)
	return args.Error(0)
}

func (e2 employeeRepoMock) Delete(id uuid.UUID) error {
	args := e2.Called(id)
	return args.Error(0)
}

func (e2 employeeRepoMock) GetById(id uuid.UUID) (*domain.Employee, error) {
	args := e2.Called(id)
	return args.Get(0).(*domain.Employee), args.Error(1)
}

type salaryRepoMock struct {
	mock.Mock
}

func (s salaryRepoMock) Save(e domain.Salary) error {
	args := s.Called(e)
	return args.Error(0)
}

func (s salaryRepoMock) Delete(id uuid.UUID) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s salaryRepoMock) GetById(id uuid.UUID) (*domain.Salary, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Salary), args.Error(1)
}
