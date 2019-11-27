package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")			//这句代码不会被执行
	os.Exit(3)
}
