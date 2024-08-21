package framework

import (
	"go-clean-arch-example/internal/application/usecases"
	"go-clean-arch-example/internal/infrastructure/inputports/http/handler"
	"go-clean-arch-example/internal/infrastructure/interfaceadapters/storage/memory"
	"net/http"
)

func NewApp() *http.ServeMux {
	mux := http.NewServeMux()

	// repositories
	repo := memory.NewInMemoryEmployeeRepository()

	// use-cases
	ucAddEmployee := usecases.InitEmployeeUseCase(repo)

	// routes
	mux.HandleFunc("POST /v1/employees", handler.AddEmployeeHandler(ucAddEmployee))

	return mux
}
