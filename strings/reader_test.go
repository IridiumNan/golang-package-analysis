package test

import (
	"io"
	_ "io"
	"os"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	r := strings.NewReader("Hello world")

	// act as the [io.WriterTo] interface
	r.WriteTo(os.Stdout)

	// NOTE: The reader has been used and you should reset it before use
	r.Reset("Hello world")

	var b strings.Builder
	_, err := io.CopyN(&b, r, 3)
	if err != nil {
		t.Errorf("error while copying 3 bytes from reader to builder, err: %s", err.Error())
	}

	expected := "Hel"

	bs := b.String()
	if bs != expected {
		t.Errorf("error when compare the builder string, expected: %s, got: %s", expected, bs)
	}
}
