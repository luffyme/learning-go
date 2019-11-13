package main

import (
	"fmt"
)

func incr(p *int) int {
    *p++
    return *p
}

func main() {
    str := "hello"
    //str[0] = 'x'
	fmt.Println(str[0])
	
	p :=1
    incr(&p)
    fmt.Println(p)
}