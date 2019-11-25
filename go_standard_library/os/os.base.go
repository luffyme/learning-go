package main

import (
    "fmt"
    "os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))

	fmt.Println(os.Getgid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
}