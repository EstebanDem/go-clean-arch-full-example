package domain

import (
	"errors"
	"github.com/google/uuid"
	"go-clean-crud-mine/internal/pkg"
	"regexp"
	"time"
)

var (
	ErrInvalidName    = errors.New("invalid employee name")
	ErrInvalidCountry = errors.New("invalid country")
)

type Employee struct {
	Id        uuid.UUID
	Name      string
	Country   string
	Salary    *Salary
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeeRepository interface {
	Save(e Employee) error
	Delete(id uuid.UUID) error
	GetById(id uuid.UUID) (*Employee, error)
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

	salary, err := NewSalary(currency, value)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Employee{
		Id:        pkg.GenerateUUID(),
		Name:      name,
		Country:   country,
		Salary:    salary,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
