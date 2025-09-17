package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/cucumber/godog"
	"github.com/oliveagle/jsonpath"
	"github.com/yogiis/golang-api-automation/helper"
)

type Entity struct {
	UrlEndpoint  string
	ResponseData *http.Response
	Cases        helper.Case
	ResponseBody []byte
}

func (e *Entity) GivenEndpoint(host, endpoint string) error {
	e.UrlEndpoint = host + endpoint

	return nil
}

func (e *Entity) SendPUTEndpointWithBodyJSON(params string, requestBody *godog.DocString) error {
	e.UrlEndpoint = e.UrlEndpoint + params

	hitEndpoint, err := http.NewRequest(http.MethodPut, e.UrlEndpoint, bytes.NewBuffer([]byte(requestBody.Content)))
	helper.LogPanicln(err)

	hitEndpoint.Header.Add("Content-Type", "application/json")

	e.ResponseData, err = sendHTTPRequest(hitEndpoint)
	helper.LogPanicln(err)

	e.ResponseBody, err = io.ReadAll(e.ResponseData.Body)
	helper.LogPanicln(err)

	defer e.ResponseData.Body.Close()

	return nil
}

func (e *Entity) ValidateStatusCode(expected int) error {
	if e.ResponseData.StatusCode != expected {
		return fmt.Errorf("expected status code %d, but got %d", expected, e.ResponseData.StatusCode)
	}

	return nil
}

var json_data map[string]interface{}

func (e *Entity) ValidateResponseBody(path, expected string) error {
	jsonpath, err := jsonpath.Compile(path)
	helper.LogPanicln(err)

	json.Unmarshal(e.ResponseBody, &json_data)
	e.assertEqualByValue(jsonpath, expected)

	return nil
}

func (e *Entity) assertEqualByValue(jsonPath *jsonpath.Compiled, expected string) {
	actual, err := jsonPath.Lookup(json_data)
	helper.LogPanicln(err)

	e.Cases.AssertEqual(expected, actual, helper.ErrorHandleEqual(expected, actual))
}

func sendHTTPRequest(hitEndpoint *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(hitEndpoint)
	helper.LogPanicln(err)

	return response, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	step := &Entity{}

	ctx.Step(`^I have an API "([^"]*)" with path "([^"]*)"$`, step.GivenEndpoint)
	ctx.Step(`^I send a PUT request with params "([^"]*)" and following body:$`, step.SendPUTEndpointWithBodyJSON)
	ctx.Step(`^status code should be (\d+)$`, step.ValidateStatusCode)
	ctx.Step(`^value "([^"]*)" should equal "([^"]*)"$`, step.ValidateResponseBody)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
