package main

import (
	"log"

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
	tz, err := tokenizer.CreateWithFile("/Users/almerlucke/Documents/Projects/Go/src/github.com/almerlucke/glisp/examples/source.glisp")
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
