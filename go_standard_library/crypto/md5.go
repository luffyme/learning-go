package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"crypto/md5"
)

func main() {
	//字符串md5
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x\n", h.Sum(nil))
	
	//文件MD5
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h5 := md5.New()
	if _, err := io.Copy(h5, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", h5.Sum(nil))
}