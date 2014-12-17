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

func (assertActual *AssertActual) Contains(expected interface{}) {
	if reflect.TypeOf(assertActual.actual).Kind() == reflect.Slice {
		actual := reflect.ValueOf(assertActual.actual)

		if find(actual, expected) == nil {
			assertActual.Test.Fail()
		}
	} else {
		assertActual.Test.Errorf("actual is incompatible with expected : %s", reflect.TypeOf(assertActual.actual).String())
	}
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

func (assertActual *AssertActual) IsNotEqualTo(expected string) {
	if assertActual.actual == expected {
		assertActual.Test.Fail()
	}
}

func (assert *Assert) That(actual interface{}) *AssertActual {
	return &AssertActual{Test: assert.Test, actual: actual}
}
