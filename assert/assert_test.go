package assert

import "testing"

func TestAssertTrue(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(true).IsTrue()

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertTrueFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(false).IsTrue()

	testResult := assert.Test.(*MockTest)
	if !testResult.Failed() || testResult.errors[0] != "Actual is not true" {
		test.FailNow()
	}
}

func TestAssertFalse(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(false).IsFalse()

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertFalseFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(true).IsFalse()

	testResult := assert.Test.(*MockTest)
	if !testResult.Failed() || testResult.errors[0] != "Actual is not false" {
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

func TestAssertEqualWithInteger(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(1).IsEqualTo(1)

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertEqualWithList(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]int{1, 2}).IsEqualTo([]int{1, 2})

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertEqualFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("Hello").IsEqualTo("different Hello")

	testResult := assert.Test.(*MockTest)
	if !testResult.Failed() || testResult.errors[0] != "Expected : \"different Hello\"\nis not Equal to:\n\"Hello\"" {
		test.FailNow()
	}
}

func TestAssertEqualFailedWithInteger(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(1).IsEqualTo(2)

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertEquaFailedlWithList(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]int{1, 2}).IsEqualTo([]int{1, 3})

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNotEqual(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("Hello").IsNotEqualTo("different Hello")

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNotEqualWithList(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]int{1, 3}).IsNotEqualTo([]int{1, 4})

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNotEqualFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("Hello").IsNotEqualTo("Hello")

	testResult := assert.Test.(*MockTest)
	if !testResult.Failed() || testResult.errors[0] != "Expected : \"Hello\"\nis Equal to:\n\"Hello\"" {
		test.FailNow()
	}
}

func TestAssertNotEqualFailedWithList(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]int{1, 3}).IsNotEqualTo([]int{1, 3})

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertContains(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]string{"hello", "world"}).Contains("world")

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertContainsInteger(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]int{3, 7}).Contains(7)

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

func TestAssertContainsIsWrongBecauseOfTypeOfActual(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(3).Contains("world!")

	testResult := assert.Test.(*MockTest)
	if !testResult.Wrong() {
		testResult.FailNow()
	}
}

func TestAssertSize(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]string{"hello", "world"}).HasSize(2)

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertSizeIsFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That([]string{"hello", "world"}).HasSize(5)

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertSizeIsWrongBecauseOfTypeOfActual(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("hello").HasSize(1)

	testResult := assert.Test.(*MockTest)
	if !testResult.Wrong() {
		test.FailNow()
	}
}

func TestAssertNil(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(nil).IsNil()

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNilIsFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("hello").IsNil()

	if !assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNotNil(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That("hello").IsNotNil()

	if assert.Test.Failed() {
		test.FailNow()
	}
}

func TestAssertNotNilIsFailed(test *testing.T) {
	assert := Assert{Test: &MockTest{}}

	assert.That(nil).IsNotNil()

	if !assert.Test.Failed() {
		test.FailNow()
	}
}
