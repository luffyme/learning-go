package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[3])

	/*
		go run os.args.go a b c d
		[/tmp/go-build549550554/b001/exe/os.args a b c d]
		[a b c d]
		c
	*/
}