package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Writer()
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read data")
}

func (f *File) Writer() {
	fmt.Println("Write data")
}

func Test(rw ReaderWriter) {
	rw.Read()
	rw.Writer()
}

func main() {
	var f File

	Test(&f)
}
