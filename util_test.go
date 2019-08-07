package utils

import (
	"testing"
)

func testFuncFormat(testStr string, expectedStr string) func(*testing.T) {
	return func(t *testing.T) {
		actual := format(testStr)
		if actual != expectedStr {
			t.Errorf("expected '%s' actual '%s'", expectedStr, actual)
		}
	  }
}

func TestFormat(t *testing.T) {
	t.Run("when '\t Hello, World\n ' is passed", testFuncFormat("\t Hello, World\n ", "Hello, World"))
	t.Run("when '     Hello, World   ' is passed", testFuncFormat("        Hello, World          ", "Hello, World"))
}

