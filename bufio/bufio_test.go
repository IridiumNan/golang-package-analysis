package test

import (
	"bufio"
	_ "io"
	"log"
	"os"
	"testing"
)

/*

To run this test file
use shell command
cd bufio && go test -v

*/

var sum = 0

const hugeFileName = "hugeFile.txt"

func doSomeWork(data []byte) {
	sum += len(data)
}

// initBigFile create a big text file
func initBigFile() error {
	hugeFile, err := os.OpenFile(hugeFileName, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("error while creating huge file, err: %s", err.Error())
	}

	defer hugeFile.Close()
	// buffer size 1 M
	w := bufio.NewWriterSize(hugeFile, 2<<20)

	for range 300000 {
		_, err = w.WriteString("This is a huge file\n")
		if err != nil {
			return err
		}
	}
	return w.Flush()
}

func createHugeFileIfNotExist(t *testing.T) {
	if _, err := os.Stat(hugeFileName); err == nil {
		return
	}

	t.Logf("hugeFile not found, try to init one, file name: %s", hugeFileName)
	err := initBigFile()
	if err != nil {
		t.Fatalf("error when init huge file: %s", err.Error())
	}
}

func TestReaderReadBytes(t *testing.T) {
	createHugeFileIfNotExist(t)
	file, err := os.Open(hugeFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		// NOTE: ReadBytes function will return the copy of this []byte slice
		line, err := r.ReadBytes('\n')
		if err != nil {
			break
		}

		doSomeWork(line)
	}

	t.Logf("The length of all byte data: %d", sum)
}

func TestReaderReadSlice(t *testing.T) {
	createHugeFileIfNotExist(t)
	file, err := os.Open(hugeFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	sum = 0
	for {
		// ReadSlice will not copy the []byte data, it return it's pointer
		line, err := r.ReadSlice('\n')
		if err != nil {
			break
		}

		doSomeWork(line)

	}

	t.Logf("The length of all byte data: %d", sum)
}

func TestScanner(t *testing.T) {
	createHugeFileIfNotExist(t)
	file, err := os.Open(hugeFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	sum = 0
	for s.Scan() {
		// the Err usually [io.EOF]
		if err := s.Err(); err != nil {
			break
		}
		sum += len(s.Bytes())
	}

	t.Logf("The length of all byte data: %d", sum)
}
