package memory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-clean-arch-example/internal/domain"
	"log"
	"testing"
	"time"
)

var salaryId = uuid.MustParse("fb20b5c4-f82c-4965-a021-3c3edd7f4af9")

func TestInMemorySalaryRepository_GetById(t *testing.T) {
	repo := buildInMemorySalaryRepo()
	salary, err := repo.GetById(salaryId)

	assert.NoError(t, err)
	assert.Equal(t, "ARS", salary.Currency)
	assert.Equal(t, 250.0, salary.Value)
}

func TestInMemorySalaryRepository_GetByIdNotFound(t *testing.T) {
	repo := buildInMemorySalaryRepo()
	_, err := repo.GetById(uuid.MustParse("fb20b5c4-f82c-4444-a021-3c3edd7f4af9"))

	assert.Equal(t, domain.ErrSalaryNotFound, err)
}

func TestInMemorySalaryRepository_Delete(t *testing.T) {
	repo := buildInMemorySalaryRepo()
	err := repo.Delete(salaryId)

	assert.NoError(t, err)
	// it should be removed
	_, err = repo.GetById(salaryId)
	assert.Equal(t, domain.ErrSalaryNotFound, err)
}

func TestInMemorySalaryRepository_DeleteNotFound(t *testing.T) {
	repo := buildInMemorySalaryRepo()
	err := repo.Delete(uuid.MustParse("fb20b5c4-f82c-4444-a021-3c3edd7f4af9"))

	assert.Equal(t, domain.ErrSalaryNotFound, err)
}

func TestInMemorySalaryRepository_Save(t *testing.T) {
	repo := buildInMemorySalaryRepo()
	err := repo.Save(domain.Salary{
		Id:        uuid.MustParse("fb20b5c4-f82c-1234-a021-3c3edd7f4af9"),
		Currency:  "ARS",
		Value:     333.33,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	assert.NoError(t, err)
	_, err = repo.GetById(salaryId)
	assert.NoError(t, err)
}

func buildInMemorySalaryRepo() InMemorySalaryRepository {
	repo := NewInMemorySalaryRepository()
	err := repo.Save(domain.Salary{
		Id:        salaryId,
		Currency:  "ARS",
		Value:     250.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Fatal("error saving in memory repo")
	}

	return repo
}
