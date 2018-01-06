package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/globals/tables"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
)

func main() {
	env := environment.New()

	fmt.Printf("\nGLISP v0.1 -- use (exit) to quit\n\n%v> ", env.CurrentNamespace().Name())

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		rd := reader.New(strings.NewReader(scanner.Text()), tables.DefaultReadTable, tables.DefaultDispatchTable, env)

		obj, err := rd.ReadObject()
		var result types.Object

		for err == nil {
			result, err = env.Eval(obj, nil)
			if err != nil {
				fmt.Printf("<! %v >\n", err)
			}

			obj, err = rd.ReadObject()
		}

		if err != nil && err != io.EOF {
			fmt.Printf("<! %v >\n", err)
		}

		if result != nil {
			fmt.Printf("%v\n", result)
		}

		fmt.Printf("%v> ", env.CurrentNamespace().Name())
	}
}
