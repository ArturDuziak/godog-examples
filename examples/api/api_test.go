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
	"github.com/stretchr/testify/assert"
)

type Entity struct {
	UrlEndpoint  string
	ResponseData *http.Response
	Cases        Case
	ResponseBody []byte
}

func (e *Entity) GivenEndpoint(host, endpoint string) error {
	e.UrlEndpoint = host + endpoint

	return nil
}

func (e *Entity) SendPUTEndpointWithBodyJSON(params string, requestBody *godog.DocString) error {
	e.UrlEndpoint = e.UrlEndpoint + params

	hitEndpoint, err := http.NewRequest(http.MethodPut, e.UrlEndpoint, bytes.NewBuffer([]byte(requestBody.Content)))
	LogPanicln(err)

	hitEndpoint.Header.Add("Content-Type", "application/json")

	e.ResponseData, err = sendHTTPRequest(hitEndpoint)
	LogPanicln(err)

	e.ResponseBody, err = io.ReadAll(e.ResponseData.Body)
	LogPanicln(err)

	defer e.ResponseData.Body.Close()

	return nil
}

func (e *Entity) SendGETEndpointWithParams(params string) error {
	e.UrlEndpoint = e.UrlEndpoint + params

	hitEndpoint, err := http.NewRequest(http.MethodGet, e.UrlEndpoint, nil)
	LogPanicln(err)

	e.ResponseData, err = sendHTTPRequest(hitEndpoint)
	LogPanicln(err)

	e.ResponseBody, err = io.ReadAll(e.ResponseData.Body)
	LogPanicln(err)

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
	LogPanicln(err)

	json.Unmarshal(e.ResponseBody, &json_data)
	e.assertEqualByValue(jsonpath, expected)

	return nil
}

func (e *Entity) ValidateResponseBodyWithJSON(expectedJSON *godog.DocString) error {
	if !assert.JSONEq(nil, expectedJSON.Content, string(e.ResponseBody)) {
		return fmt.Errorf("JSON mismatch:\nExpected: %s\nActual: %s",
			expectedJSON.Content, string(e.ResponseBody))
	}
	return nil
}

func (e *Entity) assertEqualByValue(jsonPath *jsonpath.Compiled, expected string) {
	actual, err := jsonPath.Lookup(json_data)
	LogPanicln(err)

	e.Cases.AssertEqual(expected, actual, ErrorHandleEqual(expected, actual))
}

func sendHTTPRequest(hitEndpoint *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(hitEndpoint)
	LogPanicln(err)

	return response, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	step := &Entity{}

	ctx.Step(`^I have an API "([^"]*)" with path "([^"]*)"$`, step.GivenEndpoint)
	ctx.Step(`^I send a PUT request with params "([^"]*)" and following body:$`, step.SendPUTEndpointWithBodyJSON)
	ctx.Step(`^I send a GET request with params "([^"]*)"$`, step.SendGETEndpointWithParams)
	ctx.Step(`^status code should be (\d+)$`, step.ValidateStatusCode)
	ctx.Step(`^value "([^"]*)" should equal "([^"]*)"$`, step.ValidateResponseBody)
	ctx.Step(`^response body should be:$`, step.ValidateResponseBodyWithJSON)
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
