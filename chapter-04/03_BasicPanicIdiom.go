package main

import "errors"

func main() {
	panic(errors.New("Something bad happened to make me panic. -Go Compiler"))
}
