package assert

import "testing"

func TestAssertIsTrue(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That(true).IsTrue()
}

func TestAssertIsFalse(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That(false).IsFalse()
}

func TestAssertIsEqual(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That("hello").IsEqualTo("hello")
}

func TestAssertIsNotEqual(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That("hello").IsNotEqualTo("different hello")
}

func TestAssertContainsElement(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That([]string{"hello", "world"}).Contains("world")
}

func TestAssertHasSize(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That([]string{"hello", "world"}).HasSize(2)
}

func TestAssertIsNil(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That(nil).IsNil()
}

func TestAssertIsNotNil(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That("hello").IsNotNil()
}
