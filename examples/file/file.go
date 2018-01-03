package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/globals/tables"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
)

func main() {
	f, err := os.Open("./examples/file/source.glisp")
	if err != nil {
		log.Fatal("Can't open file")
	}

	defer f.Close()

	env := environment.New()
	rd := reader.New(bufio.NewReader(f), tables.DefaultReadTable, tables.DefaultDispatchTable, env)

	obj, err := rd.ReadObject()
	var result types.Object

	for err == nil {
		result, err = env.Eval(obj, nil)
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
