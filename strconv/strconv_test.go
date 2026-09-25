package test

import (
	"strconv"
	"testing"
)

func compareAny(want, got any, t *testing.T) {
	if want != got {
		t.Errorf("err, want: %s, got: %s", want, got)
		return
	}

	t.Logf("%s = %s", want, got)
}

func TestToStr(t *testing.T) {
	compareAny("false", strconv.FormatBool(false), t)
	compareAny("true", strconv.FormatBool(true), t)

	compareAny("1", strconv.FormatInt(1, 10), t)
	compareAny("2930416610", strconv.FormatUint(2930416610, 10), t)

	// if you want to append result string into []byte, you can use Appendxxx functions
	// for example, add the boolean string
	base := []byte("allow-lan: ")
	res := strconv.AppendBool(base, false)

	compareAny("allow-lan: false", string(res), t)
}

func TestFromStr(t *testing.T) {
	// FALSE, False, false, F, f, 0 can be parsed properly
	F, err := strconv.ParseBool("F")
	if err != nil {
		t.Errorf("err when parsing boolean: %s", err.Error())
	}
	compareAny(false, F, t)

	// [strconv.Atoi] equals to [strconv.ParseInt](num, 10, 0)
	num, err := strconv.Atoi("123")
	if err != nil {
		t.Errorf("err when parsing int, err: %s", err.Error())
	}

	compareAny(123, num, t)
}
