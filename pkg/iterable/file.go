package iterable

import (
	"bufio"
	"log"
	"os"
)

type IterableFile struct {
	file    *os.File
	scanner *bufio.Scanner
}

func (f *IterableFile) Next() (value string, done bool) {
	if f.scanner.Scan() {
		return f.scanner.Text(), false
	}
	if err := f.scanner.Err(); err != nil {
		log.Fatal(err)
	}
	f.file.Close()
	return "", true
}

func NewIterableFile(input_path string) *IterableFile {
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}
	return &IterableFile{file, bufio.NewScanner(file)}
}
