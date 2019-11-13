package main

import (
	"fmt"
)

func main() {
    s := [3]int{1, 2, 3}
    a := s[:0]
    b := s[:2]
	c := s[1:2:cap(s)]
	
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	fmt.Println(len(c), cap(c))

	var m map[string]int        //A
    m["a"] = 1
    if v := m["b"]; v != nil {  //B
        fmt.Println(v)
    }
}