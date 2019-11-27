package main

import (
    "fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))

	for _, e := range os.Environ() {
		//fmt.Println(e)
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], " = ", pair[1])
	}

	fmt.Println(os.Getgid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
}