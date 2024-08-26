package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/domain"
	"time"
)

// EmployeeRecord represents each row in 'employee' table
type EmployeeRecord struct {
	id        uint64
	uuid      uuid.UUID
	Name      string
	SalaryId  uint64 //FK
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SalaryRecord represents each row in 'salary' table
type SalaryRecord struct {
	id       uint64
	Currency string
	Wage     float64
}

// EmployeeWithSalary represents the result after joining tables
type EmployeeWithSalary struct {
	id        uint64 // is it needed?
	uuid      uuid.UUID
	Name      string
	Country   string
	Currency  string
	Wage      float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeeRepositoryMySql struct {
	db *sql.DB
}

func NewMySqlEmployeeRepository() (EmployeeRepositoryMySql, error) {
	cfg := mysql.Config{
		User:   "user",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3357",
		DBName: "company_db",
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return EmployeeRepositoryMySql{}, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return EmployeeRepositoryMySql{}, pingErr
	}
	fmt.Println("Connected to MySql successfully")

	return EmployeeRepositoryMySql{
		db: db,
	}, nil

}

func (er EmployeeRepositoryMySql) Save(ctx context.Context, e domain.Employee) error {
	result, err := er.db.ExecContext(ctx, `insert into salary (currency, wage)
	values (?, ?)`, e.Salary.Currency, e.Salary.Value)
	if err != nil {
		return err
	}

	salaryId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	_, err = er.db.ExecContext(ctx, `insert into employee (uuid, name, country, salary_id, created_at, updated_at)
	values (?, ?, ?, ?, ?, ?)`, e.Id.String(), e.Name, e.Country, salaryId, e.CreatedAt, e.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (er EmployeeRepositoryMySql) Delete(ctx context.Context, id uuid.UUID) error {
	employeeRow, err := er.getRowFromUUID(ctx, id)
	if err != nil {
		return err
	}

	_, err = er.db.ExecContext(ctx, "delete from employee where id = ?", employeeRow.id)
	if err != nil {
		return err
	}

	_, err = er.db.ExecContext(ctx, "delete from salary where id = ?", employeeRow.SalaryId)
	if err != nil {
		return err
	}

	return nil
}

func (er EmployeeRepositoryMySql) GetById(ctx context.Context, id uuid.UUID) (*domain.Employee, error) {
	employeeRow, err := er.getRowFromUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	var employeeWithSalary EmployeeWithSalary
	row := er.db.QueryRowContext(ctx, `
		select e.uuid, e.name, e.country, s.currency, s.wage, e.created_at, e.updated_at
    	from employee e 
		inner join salary s on s.id = e.salary_id
		where e.id = ?
	`, employeeRow.id)

	if err := row.Scan(&employeeWithSalary.uuid,
		&employeeWithSalary.Name,
		&employeeWithSalary.Country,
		&employeeWithSalary.Currency,
		&employeeWithSalary.Wage,
		&employeeWithSalary.CreatedAt,
		&employeeWithSalary.UpdatedAt); err != nil {
		return nil, err
	}

	empDomain := toEmployeeDomain(employeeWithSalary)
	return &empDomain, nil
}

func (er EmployeeRepositoryMySql) getRowFromUUID(ctx context.Context, uuId uuid.UUID) (EmployeeRecord, error) {
	var employeeRecord EmployeeRecord
	row := er.db.QueryRowContext(ctx, "select id, salary_id from employee where uuid = ?", uuId.String())
	if err := row.Scan(&employeeRecord.id, &employeeRecord.SalaryId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return EmployeeRecord{}, fmt.Errorf("no such employee with uuid: %s", uuId.String())
		}
		return EmployeeRecord{}, err
	}
	return employeeRecord, nil
}

func toEmployeeDomain(er EmployeeWithSalary) domain.Employee {
	return domain.Employee{
		Id:      er.uuid,
		Name:    er.Name,
		Country: er.Country,
		Salary: domain.Salary{
			Currency: er.Currency,
			Value:    er.Wage,
		},
		CreatedAt: er.CreatedAt,
		UpdatedAt: er.UpdatedAt,
	}
}
