package test

import (
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	r := strings.NewReplacer("{name}", "cai", "{language}", "Chinese")

	var b strings.Builder
	tmpl := `
	Hello, I'm {name} and I speak {language}.
	Hei, family name {name} is has a long history.`

	expected := `
	Hello, I'm cai and I speak Chinese.
	Hei, family name cai is has a long history.`

	// use the Replace function
	// the old string in input args will be replaced by new strings
	// return the new replaced string directly
	res := r.Replace(tmpl)

	if res != expected {
		t.Errorf("error output, expected: %s, got: %s", expected, res)
	}

	// Use WriteString function, write the output into a [io.Writer]
	_, err := r.WriteString(&b, tmpl)
	if err != nil {
		t.Errorf("error while calling write string into builder, err: %s", err.Error())
	}

	bs := b.String()
	if bs != expected {
		t.Errorf("error output, expected: %s, got: %s", expected, bs)
	}
}
