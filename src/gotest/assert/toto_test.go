package assert

import "testing"

func TestToto(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That("toto").IsEqualTo("toto")
	assert.That("toto").IsEqualTo("toto")
}

func TestTata(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That("toto").IsEqualTo("toto")
	assert.That("toto").IsEqualTo("toto")
}
