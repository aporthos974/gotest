package assert

import (
	"reflect"
	"runtime"
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
	assertActual.verifySliceType()

	actual := reflect.ValueOf(assertActual.actual)
	if find(actual, expected) == nil {
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) HasSize(expected interface{}) {
	assertActual.verifySliceType()

	actual := reflect.ValueOf(assertActual.actual)
	if actual.Len() != expected {
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) IsNil() {
	if assertActual.actual != nil {
		assertActual.Test.Fail()
	}
}

func (assertActual *AssertActual) IsNotNil() {
	if assertActual.actual == nil {
		assertActual.Test.Fail()
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

func (assertActual *AssertActual) verifySliceType() {
	if reflect.TypeOf(assertActual.actual).Kind() != reflect.Slice {
		assertActual.errorNow("type of actual is incompatible with the expected : %s", reflect.TypeOf(assertActual.actual).String())
	}
}

func (assertActual *AssertActual) errorNow(format string, param interface{}) {
	assertActual.Test.Errorf(format, param)
	runtime.Goexit()
}
