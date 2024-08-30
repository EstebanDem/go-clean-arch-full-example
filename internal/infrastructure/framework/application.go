package framework

import (
	"go-clean-arch-example/internal/application/services"
	"go-clean-arch-example/internal/application/usecases"
	"go-clean-arch-example/internal/domain"
	"go-clean-arch-example/internal/infrastructure/inputports/http/handler"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/restclients"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/memory"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/mongodb"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/mysql"
	"log"
	"net/http"
	"os"
)

func NewApp(storage, currencyConverter string) *http.ServeMux {
	mux := http.NewServeMux()

	// repositories
	repo, err := initEmployeeRepository(storage)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	// clients
	currencyConverterClient := initCurrencyClient(currencyConverter)

	// use-cases
	ucAddEmployee := usecases.InitEmployeeUseCase(repo)
	ucGetEmployeeSalary := usecases.InitGetEmployeeSalaryUseCase(repo, currencyConverterClient)

	// routes
	mux.HandleFunc("POST /v1/employees", handler.AddEmployeeHandler(ucAddEmployee))
	mux.HandleFunc("GET /v1/employees/{id}/salary/{currency}", handler.GetEmployeeSalaryHandler(ucGetEmployeeSalary))

	return mux
}

func initEmployeeRepository(storage string) (domain.EmployeeRepository, error) {
	switch storage {
	case "mysql":
		return mysql.NewMySqlEmployeeRepository()
	case "mongo":
		return mongodb.NewEmployeeRepositoryMongo()
	default:
		return memory.NewInMemoryEmployeeRepository(), nil
	}
}

func initCurrencyClient(currencyConverter string) services.CurrencyConverter {
	switch currencyConverter {
	case "api":
		return restclients.NewFreeCurrencyApiClient(os.Getenv("API_KEY"))
	default:
		return restclients.NewPresetCurrencyConverter()
	}
}
