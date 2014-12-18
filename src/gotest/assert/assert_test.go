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

	if !assert.Test.Failed() {
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
