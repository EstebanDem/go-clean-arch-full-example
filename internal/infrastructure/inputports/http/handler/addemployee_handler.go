package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/application/usecases"
	"net/http"
	"time"
)

type AddEmployeeRequestJson struct {
	Name    string            `json:"name" validate:"required"`
	Country string            `json:"country" validate:"required"`
	Salary  SalaryRequestJson `json:"salary" validate:"required"`
}

type SalaryRequestJson struct {
	Currency string  `json:"currency" validate:"required"`
	Value    float64 `json:"value" validate:"required,gte=0"`
}

type AddEmployeeResponseJson struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func AddEmployeeHandler(uc usecases.AddEmployeeUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var addEmployeeRequest AddEmployeeRequestJson
		err := json.NewDecoder(r.Body).Decode(&addEmployeeRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonValidator := validator.New()
		err = jsonValidator.Struct(addEmployeeRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := uc.AddEmployee(addEmployeeRequest.toAddEmployeeRequest())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonResponse, err := json.Marshal(toAddEmployeeResponseJson(resp))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write(jsonResponse); err != nil {
			return
		}
	}
}

func (r AddEmployeeRequestJson) toAddEmployeeRequest() usecases.AddEmployeeRequest {
	return usecases.AddEmployeeRequest{
		Name:    r.Name,
		Country: r.Country,
		Salary: usecases.SalaryRequest{
			Currency: r.Salary.Currency,
			Value:    r.Salary.Value,
		},
	}
}

func toAddEmployeeResponseJson(r usecases.AddEmployeeResponse) AddEmployeeResponseJson {
	return AddEmployeeResponseJson{
		Id:        r.Id,
		CreatedAt: r.CreatedAt,
	}
}
