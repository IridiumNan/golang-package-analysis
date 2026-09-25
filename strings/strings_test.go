package test

import (
	"strings"
	"testing"
)

func testEqual(base, candidate string, t *testing.T, expect bool) {
	if strings.EqualFold(base, candidate) != expect {
		t.Errorf("%s and %s EqualFold return false", base, candidate)
	}
}

func TestEqualFold(t *testing.T) {
	base := "hello"

	testEqual(base, "HELLO", t, true)
	testEqual(base, "Hello", t, true)
	testEqual(base, "hELlo", t, true)

	testEqual(base, "heoll", t, false)
	testEqual(base, "LloHe", t, false)
}
