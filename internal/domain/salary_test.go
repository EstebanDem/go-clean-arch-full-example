package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSalaryWrongCurrency(t *testing.T) {
	_, err := NewSalary("-", 260)

	assert.Error(t, ErrInvalidCurrency, err)
}

func TestNewSalaryWrongValue(t *testing.T) {
	_, err := NewSalary("ARS", 0)

	assert.Error(t, ErrInvalidValue, err)
}

func TestNewSalary(t *testing.T) {
	salary, err := NewSalary("ARS", 250)

	if err != nil {
		t.Errorf("got %q, happy path expected", err.Error())
	}

	assert.Equal(t, salary.Currency, "ARS")
	assert.Equal(t, salary.Value, 250.0)
}
