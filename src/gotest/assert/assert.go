package assert

import "testing"

type Assert struct {
	Test testing.TB
}

type AssertActual struct {
	actual interface{}
	Test   testing.TB
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

func (assertActual *AssertActual) IsEqualTo(expected string) {
	if assertActual.actual != expected {
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) Contains(expected string) {
	found := false
	for _, value := range assertActual.actual.([]string) {
		if value == expected {
			found = true
		}
	}
	if !found {
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) IsNotEqualTo(expected string) {
	if assertActual.actual == expected {
		assertActual.Test.Fail()
	}
}

func (assert *Assert) That(actual interface{}) *AssertActual {
	return &AssertActual{Test: assert.Test, actual: actual}
}
