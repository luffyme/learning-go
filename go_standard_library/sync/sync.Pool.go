package main

import (
	"sync"
	"bytes"
	"fmt"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

func main() {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString("path")
	b.WriteByte('=')
	b.WriteString("/search?q=flowers")

	fmt.Println(b.Bytes())
	bufPool.Put(b)
}