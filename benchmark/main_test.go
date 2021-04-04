package main

import (
	"bufio"
	"os"
	"testing"
)

const FILE_PATH = "test.txt"
const LINES_COUNT = 100000
const TEXT = "hello\n"


// This function test operation speed of writing 100 000 lines
//to a file on our os disk.
func BenchmarkWriteFile(b *testing.B) {
	for n :=0; n < b.N; n++{
		f, err := os.Create(FILE_PATH)
		if err != nil {
			panic(err)
		}

		for i :=0; i < LINES_COUNT; i++ {
			f.WriteString(TEXT)
		}

		f.Close()
	}

}

// This function test operation speed of writing 100 000 to a file
// on our buffer memory before flushing(writing) to the disk.
func BenchmarkWriteBuffered(b *testing.B) {
	for n :=0; n < b.N; n++{
		f, err := os.Create(FILE_PATH)
		if err != nil {
			panic(err)
		}

		w := bufio.NewWriter(f)

		for i :=0; i < LINES_COUNT; i++ {
			w.WriteString(TEXT)
		}

		w.Flush()
		f.Close()
	}

}
