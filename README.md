# gotest
A light framework test.

Installation
-----------
```bash
  go get -d github.com/aporthos974/gotest
```

The test file 'assert/assert_test.go' gives all possible assertions.

Example
-----
```go
import (
  . "github.com/aporthos974/gotest/assert"
  "testing"
)

func TestIsFalse(test *testing.T) {
	var assert Assert = Assert{test}

	assert.That(false).IsFalse()
}
```
