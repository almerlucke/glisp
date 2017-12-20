package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/almerlucke/glisp"
	"github.com/almerlucke/glisp/evaluator"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
)

func main() {
	env := glisp.CreateDefaultEnvironment()

	fmt.Printf("\nGLISP v0.1 -- use (exit) to quit\n\n> ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		rd := reader.New(strings.NewReader(scanner.Text()), glisp.DefaultReadTable, glisp.DefaultDispatchTable, env)

		obj, err := rd.ReadObject()
		var result types.Object

		for err == nil {
			result, err = evaluator.Eval(obj, env)
			if err != nil {
				fmt.Printf("<! %v >\n", rd.ErrorWithError(err))
			}

			obj, err = rd.ReadObject()
		}

		if err != nil && err != io.EOF {
			fmt.Printf("<! %v >\n", rd.ErrorWithError(err))
		}

		if result != nil {
			fmt.Printf("%v\n", result)
		}

		fmt.Printf("> ")
	}
}
