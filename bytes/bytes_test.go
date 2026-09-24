package test

import (
	"bytes"
	"testing"
)

/*

To run this test file
use command
cd bytes && go test bytes -v

*/

func TestBytesSplit(t *testing.T) {
	socket := []byte("golang@gmail.com")

	sep := []byte("@")

	// use the split function to get user and domain
	parts := bytes.Split(socket, sep)

	if len(parts) != 2 {
		t.Fatalf("split length err, expected: 2, got: %d", len(parts))
	}

	if !bytes.Equal(parts[0], []byte("golang")) {
		t.Fatalf("User err, expected: %s, got: %s", "golang", string(parts[0]))
	}

	if !bytes.Equal(parts[1], []byte("gmail.com")) {
		t.Fatalf("domain err, expected: %s, got: %s", "gmail.com", string(parts[1]))
	}
}

func TestReplaceInject(t *testing.T) {
	tmpl := []byte("Hello {user}, this is programming language {lang}, {user} name, you should learn fmt.Println on {lang}")

	r1 := bytes.Replace(tmpl, []byte("{user}"), []byte("dong"), 1)
	r1 = bytes.ReplaceAll(r1, []byte("{lang}"), []byte("golang"))

	e1 := []byte("Hello dong, this is programming language golang, {user} name, you should learn fmt.Println on golang")

	if !bytes.Equal(r1, e1) {
		t.Fatalf("replace err, expected: %s, got: %s", e1, r1)
	}
}

func TestTrim(t *testing.T) {
	raw := []byte("\n\nThis is golang bytes package test\n\t")

	trimedWithoutSpace := bytes.TrimSpace(raw)

	e := []byte("This is golang bytes package test")
	if !bytes.Equal(trimedWithoutSpace, e) {
		t.Fatalf("trim test err, expected: %s, got: %s", e, trimedWithoutSpace)
	}
}
