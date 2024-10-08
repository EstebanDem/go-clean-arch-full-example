package domain

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/pkg"
	"regexp"
	"time"
)

var (
	ErrInvalidName      = errors.New("invalid employee name")
	ErrInvalidCountry   = errors.New("invalid country")
	ErrEmployeeNotFound = errors.New("employee not found")
	ErrInvalidCurrency  = errors.New("invalid currency")
	ErrInvalidValue     = errors.New("invalid value, it must be positive")
)

type Employee struct {
	Id        uuid.UUID
	Name      string
	Country   string
	Salary    Salary
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Salary struct {
	Currency string
	Value    float64
}

type EmployeeRepository interface {
	Save(ctx context.Context, e Employee) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*Employee, error)
}

func NewEmployee(name string, country string, currency string, value float64) (*Employee, error) {
	match, _ := regexp.MatchString("[a-zA-Z]{1,16}", name)
	if !match {
		return nil, ErrInvalidName
	}

	match, _ = regexp.MatchString("[a-zA-Z ]{1,20}", country)
	if !match {
		return nil, ErrInvalidCountry
	}

	match, _ = regexp.MatchString("[a-zA-Z]{3}", currency)
	if !match {
		return nil, ErrInvalidCurrency
	}

	if value <= 0 {
		return nil, ErrInvalidValue
	}

	now := time.Now()
	return &Employee{
		Id:      pkg.GenerateUUID(),
		Name:    name,
		Country: country,
		Salary: Salary{
			Currency: currency,
			Value:    value,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
