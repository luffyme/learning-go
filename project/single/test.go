package main

import (
	"fmt"
)

func main() {
	var path = "/hello/:id/*name"
	var subpath = path[10:11]
	fmt.Println(subpath)
}
