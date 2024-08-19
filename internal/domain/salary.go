package domain

import (
	"errors"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/pkg"
	"regexp"
	"time"
)

var (
	ErrInvalidCurrency = errors.New("invalid currency")
	ErrInvalidValue    = errors.New("invalid value, it must be positive")
)

type Salary struct {
	Id        uuid.UUID
	Currency  string
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SalaryRepository interface {
	Save(e Salary) error
	Delete(id uuid.UUID) error
	GetById(id uuid.UUID) (*Salary, error)
}

func NewSalary(currency string, value float64) (*Salary, error) {
	match, _ := regexp.MatchString("[a-zA-Z]{3}", currency)
	if !match {
		return nil, ErrInvalidCurrency
	}

	if value <= 0 {
		return nil, ErrInvalidValue
	}

	now := time.Now()
	return &Salary{
		Id:        pkg.GenerateUUID(),
		Currency:  currency,
		Value:     value,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
