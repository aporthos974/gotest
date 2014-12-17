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

func TestAssertTrueAsUser(test *testing.T) {
	var assert Assert = Assert{test}

	assert.IsTrue(true)
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

func TestAssertFalseAsUser(test *testing.T) {
	var assert Assert = Assert{test}

	assert.IsFalse(false)
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

func TestAssertEqualAsUser(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That("hello").IsEqualTo("hello")
}

func TestAssertNotEqual(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("Hello").IsNotEqualTo("different Hello")

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNotEqualFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("Hello").IsNotEqualTo("Hello")

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNotEqualAsUser(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That("hello").IsNotEqualTo("different hello")
}

func TestAssertContains(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]string{"hello", "world"}).Contains("world")

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertContainsIsFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]string{"hello", "world"}).Contains("world!")

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertContainsAsUser(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That([]string{"hello", "world"}).Contains("world")
}

func (test *MockTest) Fail() {
	test.failed = true
}

func (test *MockTest) Failed() bool {
	return test.failed
}
