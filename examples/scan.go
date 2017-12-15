package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/almerlucke/glisp"
	"github.com/almerlucke/glisp/evaluator"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
)

func main() {
	f, err := os.Open("./examples/source.glisp")
	if err != nil {
		log.Fatal("Can't open file")
	}

	defer f.Close()

	env := glisp.CreateDefaultEnvironment()
	reader := reader.New(bufio.NewReader(f), env)

	obj, err := reader.Read()
	var result types.Object

	for err == nil {
		result, err = evaluator.Eval(obj, env)
		if err != nil {
			log.Fatalf("eval error %v\n", reader.ErrorWithError(err))
		}

		obj, err = reader.Read()
	}

	if err != nil && err != io.EOF {
		log.Fatalf("Reader err: %v\n", reader.ErrorWithError(err))
	}

	if result != nil {
		log.Printf("%v\n", result)
	}
}
