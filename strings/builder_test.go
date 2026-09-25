package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func printCapAndLen(b *strings.Builder) {
	fmt.Printf("the capacity of builder: %d, the length of builder: %d\n", b.Cap(), b.Len())
}

func TestBuilder(t *testing.T) {
	var builder strings.Builder

	printCapAndLen(&builder)

	// pre allocate capacity for this builder
	builder.Grow(1 >> 6)

	printCapAndLen(&builder)

	// write string into this builder
	builder.WriteString("Hello world")

	builder.WriteString("\nThis is Golang programming language")

	buf := bytes.NewBufferString("\nThis is Linux x86_64 platform")

	// use [bytes.Buffer.WriteTo] to write string into builder

	_, err := buf.WriteTo(&builder)
	if err != nil {
		t.Fatalf("error while writing buffered string into builder, err: %s", err.Error())
	}

	printCapAndLen(&builder)

	// Check if the result match expected string
	expected := "Hello world\nThis is Golang programming language\nThis is Linux x86_64 platform"
	bs := builder.String()
	if bs != expected {
		t.Errorf("error while building string, expected: %s, got: %s", expected, bs)
	}
}
