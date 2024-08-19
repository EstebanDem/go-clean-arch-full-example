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
	SalaryId  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeeRepository interface {
	Save(e Employee) error
	Delete(id uuid.UUID) error
	GetById(id uuid.UUID) (*Employee, error)
}

func NewEmployee(name string, country string, salaryId uuid.UUID) (*Employee, error) {
	match, _ := regexp.MatchString("[a-zA-Z]{1,16}", name)
	if !match {
		return nil, ErrInvalidName
	}

	match, _ = regexp.MatchString("[a-zA-Z ]{1,20}", country)
	if !match {
		return nil, ErrInvalidCountry
	}

	now := time.Now()
	return &Employee{
		Id:        pkg.GenerateUUID(),
		Name:      name,
		Country:   country,
		SalaryId:  salaryId,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
