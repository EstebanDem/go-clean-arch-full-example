package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-clean-arch-example/internal/application/usecases"
	"net/http"
)

type GetEmployeeSalaryJsonResponse struct {
	EmployeeId      uuid.UUID       `json:"employee_id"`
	Salary          Salary          `json:"salary"`
	ConvertedSalary ConvertedSalary `json:"converted_salary"`
}

type Salary struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type ConvertedSalary struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
	Rate     float64 `json:"rate"`
}

func GetEmployeeSalaryHandler(uc usecases.GetEmployeeSalaryUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		currency := r.PathValue("currency")
		if id == "" || currency == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := uuid.Validate(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		employeeRequest := usecases.GetEmployeeSalaryRequest{
			EmployeeId: uuid.MustParse(id),
			Currency:   currency,
		}

		employeeResponse, err := uc.GetSalary(employeeRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonResponse, err := json.Marshal(toGetEmployeeSalaryJsonResponse(employeeResponse))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(jsonResponse); err != nil {
			return
		}

	}
}

func toGetEmployeeSalaryJsonResponse(er usecases.GetEmployeeSalaryResponse) GetEmployeeSalaryJsonResponse {
	return GetEmployeeSalaryJsonResponse{
		EmployeeId: er.EmployeeId,
		Salary: Salary{
			Currency: er.Salary.Currency,
			Value:    er.Salary.Value,
		},
		ConvertedSalary: ConvertedSalary{
			Currency: er.ConvertedSalary.Currency,
			Value:    er.ConvertedSalary.Value,
			Rate:     er.ConvertedSalary.Rate,
		},
	}
}
