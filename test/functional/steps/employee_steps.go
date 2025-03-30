package steps

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"net/http"
	"os"
	"reflect"
)

var (
	client       *http.Client
	url          string
	statusCode   int
	employeeId   uuid.UUID
	currencyReq  string
	responseBody map[string]interface{}
)

func anEmployeeWithIdAndSalaryIn(id, currency string) {
	employeeId = uuid.MustParse(id)
	currencyReq = currency
}

func theirInformationIsRequested() error {
	url = fmt.Sprintf("http://localhost:1234/v1/employees/%s/salary/%s", employeeId, currencyReq)
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error making the request: %v", err)
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return fmt.Errorf("error decoding response body: %v", err)
	}

	return nil
}

func theResponseStatusCodeShouldBe(expected int) error {
	if statusCode != expected {
		return fmt.Errorf("expected status code %d, got %d", expected, statusCode)
	}
	return nil
}
func theResponseBodyShouldBe(fileName string) error {
	jsonPath := fileName + ".json"
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("error reading json file: %s, error: %s", jsonPath, err)
	}

	var fromFile map[string]interface{}
	err = json.Unmarshal(file, &fromFile)
	if err != nil {
		return fmt.Errorf("error unmarshalling file, %s", err)
	}

	if reflect.DeepEqual(fromFile, responseBody) {
		return nil
	} else {
		fmt.Printf("Response: %+v\n", responseBody)
		fmt.Printf("Expected: %+v\n", fromFile)
		return fmt.Errorf("reponses are different")
	}

}

func InitializeScenario(ctx *godog.ScenarioContext) {
	client = &http.Client{}
	ctx.Step(`^an employee with id (\S+) and salary in ([A-Z]+)$`, anEmployeeWithIdAndSalaryIn)
	ctx.Step(`^their information is requested$`, theirInformationIsRequested)
	ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
	ctx.Step(`^the response body should be (\S+)$`, theResponseBodyShouldBe)
}
