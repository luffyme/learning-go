package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	str := "sha1 is string"
	h := sha1.New()

	h.Write([]byte(str))
	bs := h.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", bs)
}