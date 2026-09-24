package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

const targetUrl = "https://example.com"

// TestBytesBufferFprint use [bytes.Buffer] as [io.Reader] and [io.Writer]
func TestBytesBufferFprint(t *testing.T) {
	// init with Buffer.buf = nil
	// Don't worry about it
	// The internal grow function will call make function before visit this buf
	// You can always initialize a new [bytes.Buffer] if you have no init data
	var buf bytes.Buffer

	// Grow the capacity to 1 KiB
	// buf capacity will be 1 KiB
	buf.Grow(1 << 10)

	if buf.Cap() != 1024 {
		t.Errorf("grow err, expected capacity: %d, got: %d", 1024, buf.Cap())
	}

	// use the [fmt.Fprint] function which create []byte then call write function of [io.Writer]
	// In this case, the buf pointer act as [io.Writer]
	_, err := fmt.Fprint(&buf, "hello world, hello linux, hello ubuntu, hello dong\n")
	if err != nil {
		t.Errorf("error while writing hello sentence into buf, err: %s", err.Error())
	}

	// Use copy function write all data on [io.Reader] to [io.Writer]
	// In this case, buf pointer act as [io.Reader]
	_, err = io.Copy(os.Stdout, &buf)
	if err != nil {
		t.Errorf("error while reading data from buffer then copy it into stdout, err: %s", err.Error())
	}
}

// TestBytesBufferReadFrom use [bytes.Buffer] as [io.ReaderFrom] as [io.WriterTo]
func TestBytesBufferReadFrom(t *testing.T) {
	var buf bytes.Buffer

	resp, err := http.Get(targetUrl)
	if err != nil {
		t.Errorf("error while request for url: %s, err: %s", targetUrl, err.Error())
	}

	// the [http.Response.Body] is interface [io.ReadCloser]
	defer resp.Body.Close()

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Errorf("error while reading from resp body, err: %s", err.Error())
	}

	dstFile, err := os.OpenFile("example.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		t.Errorf("error while creating a dst file, filename: %s, err: %s", "example.html", err.Error())
	}

	defer dstFile.Close()

	n, err := buf.WriteTo(dstFile)
	if err != nil {
		t.Errorf("error while write buf data into file: %s, err: %s", dstFile.Name(), err.Error())
	}

	t.Logf("write the example.com body into file: %s, written size: %d", dstFile.Name(), n)
}
