package api

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	testing testing.T
}

func (t *Case) assertNew() *assert.Assertions {
	return assert.New(&t.testing)
}

func (t *Case) AssertEqual(expected, actual, err interface{}) bool {
	equal := t.assertNew().Equal(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), err)
	if !equal {
		LogPanicln(ErrorHandleEqual(expected, actual))
	}
	return true
}

func LogPanicln(err interface{}) error {
	if err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

func ErrorHandleEqual(expected, actual interface{}) string {
	return fmt.Sprintf("Expected : %v Actual : %v", expected, actual)
}
