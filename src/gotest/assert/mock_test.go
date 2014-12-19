package assert

import "testing"

type MockTest struct {
	testing.TB
	failed bool
}

func (test *MockTest) Errorf(format string, args ...interface{}) {
	test.failed = true
}

func (test *MockTest) Error(args ...interface{}) {
	test.failed = true
}

func (test *MockTest) Failed() bool {
	return test.failed
}

func (test *MockTest) Wrong() bool {
	return test.failed
}
