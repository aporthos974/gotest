package assert

import "testing"

type MockTest struct {
	testing.TB
	failed bool
	wrong  bool
}

func (test *MockTest) Fail() {
	test.failed = true
}

func (test *MockTest) Errorf(format string, args ...interface{}) {
	test.wrong = true
}

func (test *MockTest) Failed() bool {
	return test.failed
}

func (test *MockTest) Wrong() bool {
	return test.wrong
}
