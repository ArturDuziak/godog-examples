package repeatedcharacters

import (
	"context"
	"flag"
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

var (
	tags   = flag.String("godog.tags", "", "tags to execute")
	format = flag.String("godog.format", "pretty", "format")
)

var opts = &godog.Options{
	Paths: []string{"features"},
}

func TestFeatures(t *testing.T) {
	opts.TestingT = t
	opts.Tags = *tags
	opts.Format = *format

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             opts,
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
