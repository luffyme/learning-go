package main

import (
	"fmt"
)

func main() {
	append1()
	append2()
}

func append1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) 			//[0 0 0 0 0 1 2 3]
}

func append2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s) 			//[1 2 3]
}