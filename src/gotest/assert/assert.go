package assert

import "testing"

type Assert struct {
	Test testing.TB
}

type AssertExpected struct {
	expected string
	Test     testing.TB
}

func (assert *Assert) IsTrue(actual bool) {
	if !actual {
		assert.Test.Fail()
	}
}

func (assert *Assert) IsFalse(actual bool) {
	if actual {
		assert.Test.Fail()
	}
}

func (assertExpected *AssertExpected) IsEqualTo(actual string) {
	if assertExpected.expected != actual {
		assertExpected.Test.Fail()
	}
}

func (assert *Assert) That(expected string) *AssertExpected {
	return &AssertExpected{Test: assert.Test, expected: expected}
}
