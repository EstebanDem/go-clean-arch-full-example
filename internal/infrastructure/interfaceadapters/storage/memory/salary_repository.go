package memory

import (
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
	"time"
)

type InMemorySalary struct {
	Id        uuid.UUID
	Currency  string
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s InMemorySalary) toDomainSalary() domain.Salary {
	return domain.Salary{
		Id:        s.Id,
		Currency:  s.Currency,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func toInMemorySalary(s domain.Salary) InMemorySalary {
	return InMemorySalary{
		Id:        s.Id,
		Currency:  s.Currency,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

type InMemorySalaryRepository struct {
	salaries map[uuid.UUID]InMemorySalary
}

func NewInMemorySalaryRepository() InMemorySalaryRepository {
	salaries := make(map[uuid.UUID]InMemorySalary)
	return InMemorySalaryRepository{salaries: salaries}
}

func (i InMemorySalaryRepository) Save(e domain.Salary) error {
	salary := toInMemorySalary(e)
	i.salaries[salary.Id] = salary
	return nil
}

func (i InMemorySalaryRepository) Delete(id uuid.UUID) error {
	_, ok := i.salaries[id]
	if !ok {
		return domain.ErrSalaryNotFound
	}

	delete(i.salaries, id)
	return nil
}

func (i InMemorySalaryRepository) GetById(id uuid.UUID) (*domain.Salary, error) {
	salary, ok := i.salaries[id]
	if !ok {
		return nil, domain.ErrSalaryNotFound
	}
	dSalary := salary.toDomainSalary()
	return &dSalary, nil
}
