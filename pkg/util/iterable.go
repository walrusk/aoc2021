package util

import (
	"bufio"
	"log"
	"os"
)

type stringIterator interface {
	Next() (value string, done bool)
}

type iterableFile struct {
	file    *os.File
	scanner *bufio.Scanner
}

func (f *iterableFile) Next() (value string, done bool) {
	if f.scanner.Scan() {
		return f.scanner.Text(), false
	}
	if err := f.scanner.Err(); err != nil {
		log.Fatal(err)
	}
	f.file.Close()
	return "", true
}

func NewIterableFile(input_path string) *iterableFile {
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}
	return &iterableFile{file, bufio.NewScanner(file)}
}
