package main

import (
	_ "fmt"
	"os"
	"syscall"
    "os/exec"
)

func main() {
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	/*
	cmd := exec.Command("ping", "127.0.0.1")
	out, err := cmd.Output()
	if err!=nil {
		println("Command Error!", err.Error())
		return
	}
	fmt.Println(string(out))
	*/
}