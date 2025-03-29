package framework

import (
	"go-clean-arch-example/internal/application/currency"
	appEmployee "go-clean-arch-example/internal/application/employee"
	"go-clean-arch-example/internal/domain/employee"
	emplHandler "go-clean-arch-example/internal/infrastructure/inputports/http/employee"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/restclients/currencyconverter"
	memoryemplrepo "go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/memory/employee"
	mongoemplrepo "go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/mongodb/employee"
	mysqlemplrepo "go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/mysql/employee"
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
	ucAddEmployee := appEmployee.InitEmployeeUseCase(repo)
	ucGetEmployeeSalary := appEmployee.InitGetEmployeeSalaryUseCase(repo, currencyConverterClient)

	// routes
	mux.HandleFunc("POST /v1/employees", emplHandler.AddEmployeeHandler(ucAddEmployee))
	mux.HandleFunc("GET /v1/employees/{id}/salary/{currency}", emplHandler.GetEmployeeSalaryHandler(ucGetEmployeeSalary))

	return mux
}

func initEmployeeRepository(storage string) (employee.EmployeeRepository, error) {
	switch storage {
	case "mysql":
		return mysqlemplrepo.NewMySqlEmployeeRepository()
	case "mongo":
		return mongoemplrepo.NewEmployeeRepositoryMongo()
	default:
		return memoryemplrepo.NewInMemoryEmployeeRepository(), nil
	}
}

func initCurrencyClient(currencyConverter string) currency.CurrencyConverter {
	switch currencyConverter {
	case "api":
		return currencyconverter.NewFreeCurrencyApiClient(os.Getenv("API_KEY"))
	default:
		return currencyconverter.NewPresetCurrencyConverter()
	}
}
