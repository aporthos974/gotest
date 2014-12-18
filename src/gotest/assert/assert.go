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
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) IsFalse() {
	if assertActual.actual.(bool) {
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) IsEqualTo(expected string) {
	if assertActual.actual != expected {
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) IsNotEqualTo(expected string) {
	if assertActual.actual == expected {
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
		assertActual.Test.Errorf("type of actual is incompatible with the expected : %s", reflect.TypeOf(assertActual.actual).String())
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
