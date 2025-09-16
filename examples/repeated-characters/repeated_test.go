package repeatedcharacters

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

// testContext holds the state between test steps
type testContext struct {
	inputWord string
	result    string
}

func (tc *testContext) givenWord(word string) error {
	tc.inputWord = word

	return nil
}

func (tc *testContext) iCountTheLetters() error {
	tc.result = RepeatedCharacters(tc.inputWord)

	return nil
}

func (tc *testContext) iShouldGet(expected string) error {
	if tc.result != expected {
		return fmt.Errorf("expected %s, but got %s", expected, tc.result)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := &testContext{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		tc.inputWord = ""
		tc.result = ""

		return ctx, nil
	})

	ctx.Step(`^the word "([^"]*)"$`, tc.givenWord)
	ctx.Step(`^I count the letters$`, tc.iCountTheLetters)
	ctx.Step(`^I should get "([^"]*)"$`, tc.iShouldGet)
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
