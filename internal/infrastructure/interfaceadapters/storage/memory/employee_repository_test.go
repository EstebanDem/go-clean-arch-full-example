package memory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-clean-arch-example/internal/domain"
	"log"
	"testing"
	"time"
)

var employeeId = uuid.MustParse("fb20b5c4-f82c-9876-a021-3c3edd7f4af9")

func TestInMemoryEmployeeRepository_GetById(t *testing.T) {
	employeeRepo := buildInMemoryEmployeeRepo()
	employee, err := employeeRepo.GetById(employeeId)

	assert.NoError(t, err)
	assert.Equal(t, "Walter", employee.Name)
	assert.Equal(t, "Argentina", employee.Country)
	assert.Equal(t, "ARS", employee.Salary.Currency)
	assert.Equal(t, 270.00, employee.Salary.Value)
}

func TestInMemoryEmployeeRepository_GetByIdNotFound(t *testing.T) {
	employeeRepo := buildInMemoryEmployeeRepo()
	_, err := employeeRepo.GetById(uuid.MustParse("fb20b5c9-f99c-9876-a021-3c3edd7f4af9"))

	assert.Equal(t, domain.ErrEmployeeNotFound, err)
}

func TestInMemoryEmployeeRepository_DeleteNotFound(t *testing.T) {
	employeeRepo := buildInMemoryEmployeeRepo()
	err := employeeRepo.Delete(uuid.MustParse("fb20b5c9-f99c-9876-a021-3c3edd7f4af9"))

	assert.Equal(t, domain.ErrEmployeeNotFound, err)
}

func TestInMemoryEmployeeRepository_Delete(t *testing.T) {
	employeeRepo := buildInMemoryEmployeeRepo()
	err := employeeRepo.Delete(employeeId)

	assert.NoError(t, err)
}

func TestInMemoryEmployeeRepository_Save(t *testing.T) {
	employeeRepo := buildInMemoryEmployeeRepo()
	err := employeeRepo.Save(domain.Employee{
		Id:      uuid.MustParse("fb20b5c9-f99c-9876-a021-3c3edd7f4af9"),
		Name:    "Hector",
		Country: "USA",
		Salary: domain.Salary{
			Currency: "USD",
			Value:    250.0,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.NoError(t, err)

	_, err = employeeRepo.GetById(uuid.MustParse("fb20b5c9-f99c-9876-a021-3c3edd7f4af9"))
	assert.NoError(t, err)
}

func buildInMemoryEmployeeRepo() InMemoryEmployeeRepository {
	employeeRepo := NewInMemoryEmployeeRepository()
	err := employeeRepo.Save(domain.Employee{
		Id:      employeeId,
		Name:    "Walter",
		Country: "Argentina",
		Salary: domain.Salary{
			Currency: "ARS",
			Value:    270.00,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Fatal("error saving in employee repo")
	}
	return employeeRepo
}
