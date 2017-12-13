package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/reader"
)

func main() {
	f, err := os.Open("./examples/source.glisp")
	if err != nil {
		log.Fatal("Can't open file")
	}

	defer f.Close()

	environment := environment.New()
	reader := reader.New(bufio.NewReader(f), environment)

	obj, err := reader.Read()
	for err == nil {
		log.Printf("obj %v\n", obj)

		obj, err = reader.Read()
	}

	if err != nil && err != io.EOF {
		log.Fatalf("Reader err: %v\n", err)
	}
}
