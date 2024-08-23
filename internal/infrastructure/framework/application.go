package framework

import (
	"go-clean-arch-example/internal/application/usecases"
	"go-clean-arch-example/internal/infrastructure/inputports/http/handler"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/restclients"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/memory"
	"net/http"
	"os"
)

func NewApp() *http.ServeMux {
	mux := http.NewServeMux()

	// repositories
	repo := memory.NewInMemoryEmployeeRepository()

	// clients
	freeCurrencyApiClient := restclients.NewFreeCurrencyApiClient(os.Getenv("API_KEY"))

	// use-cases
	ucAddEmployee := usecases.InitEmployeeUseCase(repo)
	ucGetEmployeeSalary := usecases.InitGetEmployeeSalaryUseCase(repo, freeCurrencyApiClient)

	// routes
	mux.HandleFunc("POST /v1/employees", handler.AddEmployeeHandler(ucAddEmployee))
	mux.HandleFunc("GET /v1/employees/{id}/salary/{currency}", handler.GetEmployeeSalaryHandler(ucGetEmployeeSalary))

	return mux
}
