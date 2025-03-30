package functional

import (
	"github.com/cucumber/godog"
	"go-clean-arch-example/internal/infrastructure/framework"
	"go-clean-arch-example/test/functional/steps"
	"net/http"
	"testing"
)

func startApp() {
	go func() {
		app := framework.NewApp("memory", "local")
		err := http.ListenAndServe(":1234", app)
		if err != nil {
			panic("Error serving server")
		}
	}()

}

func TestFeatures(t *testing.T) {
	startApp()
	opts := godog.Options{
		Format: "progress",
		Paths:  []string{"../../test/functional/features/employee.feature"},
	}

	status := godog.TestSuite{
		ScenarioInitializer: steps.InitializeScenario,
		Options:             &opts,
	}.Run()

	if status != 0 {
		t.Fatalf("Godog test failed with status: %d", status)
	}
}
