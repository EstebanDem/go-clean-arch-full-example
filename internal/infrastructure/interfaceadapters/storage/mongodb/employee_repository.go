package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type EmployeeRepositoryMongo struct {
	collection *mongo.Collection
}

func NewEmployeeRepositoryMongo() (EmployeeRepositoryMongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27022/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return EmployeeRepositoryMongo{}, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return EmployeeRepositoryMongo{}, err
	}
	coll := client.Database("company_db").Collection("employees")
	fmt.Println("Connected to MongoDB successfully")
	return EmployeeRepositoryMongo{
		collection: coll,
	}, nil
}

func (empRepo EmployeeRepositoryMongo) Save(ctx context.Context, e domain.Employee) error {
	doc := toEmployeeDocument(e)
	_, err := empRepo.collection.InsertOne(ctx, doc)

	return err
}

func (empRepo EmployeeRepositoryMongo) Delete(ctx context.Context, id uuid.UUID) error {
	doc, err := empRepo.getDocumentFromUUID(ctx, id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": doc.ID}
	_, err = empRepo.collection.DeleteOne(ctx, filter)
	return err
}

func (empRepo EmployeeRepositoryMongo) GetById(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	doc, err := empRepo.getDocumentFromUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	emp := toDomainEmployee(doc)
	return &emp, nil
}

func (empRepo EmployeeRepositoryMongo) getDocumentFromUUID(ctx context.Context, uuidVal uuid.UUID) (EmployeeMongoDocument, error) {
	filter := bson.M{"uuid": uuidVal}
	var doc EmployeeMongoDocument

	err := empRepo.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return EmployeeMongoDocument{}, nil
		}
		return EmployeeMongoDocument{}, err
	}

	return doc, nil
}
