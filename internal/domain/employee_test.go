package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmployeeWrongName(t *testing.T) {
	_, err := NewEmployee("123", "Argentina", "ARS", 200.50)

	assert.Equal(t, ErrInvalidName, err)
}

func TestNewEmployeeWrongCountry(t *testing.T) {
	_, err := NewEmployee("Pedro", "333", "ARS", 200.50)

	assert.Equal(t, ErrInvalidCountry, err)
}

func TestNewEmployeeWrongSalary(t *testing.T) {
	_, err := NewEmployee("Pedro", "Argentina", "ARS", 0)

	assert.Equal(t, ErrInvalidValue, err)
}

func TestNewEmployee(t *testing.T) {
	employee, err := NewEmployee("Pedro", "Argentina", "ARS", 200.50)

	assert.Nil(t, err)
	assert.Equal(t, employee.Name, "Pedro")
	assert.Equal(t, employee.Country, "Argentina")
	assert.Equal(t, employee.Salary.Currency, "ARS")
	assert.Equal(t, employee.Salary.Value, 200.50)
}
