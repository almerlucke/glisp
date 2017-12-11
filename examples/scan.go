package main

import (
	"log"
	"os"

	"github.com/almerlucke/glisp/tokenizer"
)

/*
tz, err := tokenizer.CreateWithFile("/Users/almerlucke/Documents/Projects/Go/src/github.com/almerlucke/glisp/examples/source.glisp")
if err != nil {
	log.Printf("err %v\n", err)
	return
}
*/

func main() {
	d, err := os.Getwd()
	log.Printf("wdir %s %v\n", d, err)

	tz, err := tokenizer.CreateWithFile("./examples/source.glisp")
	if err != nil {
		log.Printf("err %v\n", err)
		return
	}

	defer tz.Destroy()

	token, err := tz.NextToken()
	if err != nil {
		log.Printf("err %v\n", err)
		return
	}

	log.Printf("%v\n", token)
}
