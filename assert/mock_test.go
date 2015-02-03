package assert

import (
	"fmt"
	"testing"
)

type MockTest struct {
	testing.TB
	failed bool
	errors []interface{}
}

func (test *MockTest) Errorf(format string, args ...interface{}) {
	test.failed = true
	test.errors = append(test.errors, fmt.Sprintf(format, args...))
}

func (test *MockTest) Error(args ...interface{}) {
	test.failed = true
	for _, err := range args {
		test.errors = append(test.errors, err.(string))
	}
}

func (test *MockTest) Failed() bool {
	return test.failed
}

func (test *MockTest) Wrong() bool {
	return test.failed
}
