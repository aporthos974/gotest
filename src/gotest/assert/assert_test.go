package assert

import (
	"testing"
)

type MockTest struct {
	testing.TB
	failed bool
}

func TestAssertTrue(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.IsTrue(true)

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertTrueFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.IsTrue(false)

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertFalse(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.IsFalse(false)

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertFalseFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.IsFalse(true)

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertEqual(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("Hello").IsEqualTo("Hello")

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertEqualFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("Hello").IsEqualTo("different Hello")

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func (test *MockTest) Fail() {
	test.failed = true
}

func (test *MockTest) Failed() bool {
	return test.failed
}
