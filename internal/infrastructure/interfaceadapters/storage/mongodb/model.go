package mongodb

import (
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type EmployeeMongoDocument struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UUID      uuid.UUID          `bson:"uuid"`
	Name      string             `bson:"name"`
	Country   string             `bson:"country"`
	Salary    SalaryMongo        `bson:"salary"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type SalaryMongo struct {
	Currency string  `bson:"currency"`
	Value    float64 `bson:"value"`
}

func toDomainEmployee(e EmployeeMongoDocument) domain.Employee {
	return domain.Employee{
		Id:      e.UUID,
		Name:    e.Name,
		Country: e.Country,
		Salary: domain.Salary{
			Currency: e.Salary.Currency,
			Value:    e.Salary.Value,
		},
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func toEmployeeDocument(entity domain.Employee) EmployeeMongoDocument {
	return EmployeeMongoDocument{
		UUID:    entity.Id,
		Name:    entity.Name,
		Country: entity.Country,
		Salary: SalaryMongo{
			Currency: entity.Salary.Currency,
			Value:    entity.Salary.Value,
		},
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
