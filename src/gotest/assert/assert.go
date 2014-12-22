package assert

import (
	"reflect"
	"testing"
)

type Assert struct {
	Test testing.TB
}

type AssertActual struct {
	actual interface{}
	Test   testing.TB
}

func (assert *Assert) That(actual interface{}) *AssertActual {
	return &AssertActual{Test: assert.Test, actual: actual}
}

func (assertActual *AssertActual) IsTrue() {
	if !assertActual.actual.(bool) {
		assertActual.Test.Error("Actual is not true")
	}
}

func (assertActual *AssertActual) IsFalse() {
	if assertActual.actual.(bool) {
		assertActual.Test.Error("Actual is not false")
	}
}

func (assertActual *AssertActual) IsEqualTo(expected interface{}) {
	if !reflect.DeepEqual(assertActual.actual, expected) {
		assertActual.Test.Errorf("Expected : \"%s\"\nis not Equal to:\n\"%s\"", expected, assertActual.actual)
	}
}

func (assertActual *AssertActual) IsNotEqualTo(expected interface{}) {
	if reflect.DeepEqual(assertActual.actual, expected) {
		assertActual.Test.Errorf("Expected : \"%s\"\nis Equal to:\n\"%s\"", expected, assertActual.actual)
	}
}

func (assertActual *AssertActual) Contains(expected interface{}) {
	if !assertActual.isSliceType() {
		return
	}

	actual := reflect.ValueOf(assertActual.actual)
	if find(actual, expected) == nil {
		assertActual.Test.Errorf("actual doesn't contain %s", expected)
	}
}

func (assertActual *AssertActual) HasSize(expected interface{}) {
	if !assertActual.isSliceType() {
		return
	}

	actual := reflect.ValueOf(assertActual.actual)
	if actual.Len() != expected {
		assertActual.Test.Errorf("size of expected is %d", actual)
	}
}

func (assertActual *AssertActual) IsNil() {
	if assertActual.actual != nil {
		assertActual.Test.Error("expected isn't nil")
	}
}

func (assertActual *AssertActual) IsNotNil() {
	if assertActual.actual == nil {
		assertActual.Test.Error("expected is nil")
	}
}

func (assertActual *AssertActual) isSliceType() bool {
	if reflect.TypeOf(assertActual.actual).Kind() != reflect.Slice {
		assertActual.errorNow("type of actual is incompatible with the expected : %s", reflect.TypeOf(assertActual.actual).String())
		return false
	}
	return true
}

func (assertActual *AssertActual) errorNow(format string, param interface{}) {
	assertActual.Test.Errorf(format, param)
}

func find(elements reflect.Value, elementToFind interface{}) interface{} {
	for i := 0; i < elements.Len(); i++ {
		value := elements.Index(i).Interface()
		if value == elementToFind {
			return value
		}
	}
	return nil
}
