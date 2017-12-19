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
	rd := reader.New(bufio.NewReader(f), glisp.DefaultReadTable, glisp.DefaultDispatchTable, env)

	obj, err := rd.ReadObject()
	var result types.Object

	for err == nil {
		result, err = evaluator.Eval(obj, env)
		if err != nil {
			log.Fatalf("eval error %v\n", rd.ErrorWithError(err))
		}

		obj, err = rd.ReadObject()
	}

	if err != nil && err != io.EOF {
		log.Fatalf("Reader err: %v\n", rd.ErrorWithError(err))
	}

	if result != nil {
		log.Printf("%v\n", result)
	}
}
