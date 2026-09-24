package test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	// Wrap raw []byte with [bytes.Reader]
	r := bytes.NewReader([]byte("hello world"))

	// ReadAll function accept [io.Reader]
	data, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("error when read data from bytes.Reader, err: %s", err.Error())
	}

	var buf bytes.Buffer
	_, err = fmt.Fprint(&buf, string(data))
	if err != nil {
		t.Errorf("error when write string data into buffer, err: %s", err.Error())
	}

	if !bytes.Equal(buf.Bytes(), data) {
		t.Errorf("error when compare original data and buffer bytes, expected: %s, got: %s", buf.String(), string(data))
	}
}
