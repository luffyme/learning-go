package main

import (
	"sync"
	"time"
	"fmt"
)

var bytePool = sync.Pool {
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	a := time.Now().Unix()

	for i := 0; i <  1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}

	b := time.Now().Unix()

	for i := 0; i <  1000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}

	c := time.Now().Unix()

	fmt.Println("without pool ", b - a, "s")
	fmt.Println("with    pool ", c - b, "s")
}