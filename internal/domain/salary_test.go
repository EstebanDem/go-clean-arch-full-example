package domain

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSalaryWrongCurrency(t *testing.T) {
	_, err := NewSalary("-", 260)

	if !errors.Is(ErrInvalidCurrency, err) {
		t.Errorf("got %q, expected %q", err.Error(), ErrInvalidCurrency.Error())
	}
}

func TestNewSalaryWrongValue(t *testing.T) {
	_, err := NewSalary("ARS", 0)

	if !errors.Is(ErrInvalidValue, err) {
		t.Errorf("got %q, expected %q", err.Error(), ErrInvalidValue.Error())
	}
}

func TestNewSalary(t *testing.T) {
	salary, err := NewSalary("ARS", 250)

	if err != nil {
		t.Errorf("got %q, happy path expected", err.Error())
	}

	assert.Equal(t, salary.Currency, "ARS")
	assert.Equal(t, salary.Value, 250.0)
}
