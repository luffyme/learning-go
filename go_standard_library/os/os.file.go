package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("/tmp/data")
	if err != nil {
		fmt.Println("ReadFile error : ", err)
	}
	fmt.Println(string(data))

	f, err := os.Open("/tmp/data")
	if err != nil {
		fmt.Println("Open file error : ", err)
	}

	b := make([]byte, 5)
	n, err := f.Read(b)
	if err != nil {
		fmt.Println("Read error : ", err)
	}
	fmt.Printf("%d bytes: %s\n", n, string(b))



	s1, err := f.Seek(6, 0)
	if err != nil {
		fmt.Println("Seek file error : ", err)
	}
	sb1 := make([]byte, 2)
	n2, err := f.Read(sb1)
	if err != nil {
		fmt.Println("Seek Read file error : ", err)
	}
	fmt.Printf("%d bytes @ %d: %s\n", n2, s1, string(sb1))



    sb2, err := f.Seek(6, 0)
	if err != nil {
		fmt.Println("Seek file error : ", err)
	}
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
	if err != nil {
		fmt.Println("Seek ReadAtLeast file error : ", err)
	}
	fmt.Printf("%d bytes @ %d: %s\n", n3, sb2, string(b3))
	


	_, err = f.Seek(0, 0)
	if err != nil {
		fmt.Println("Seek file error : ", err)
	}



	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	if err != nil {
		fmt.Println("Peek file error : ", err)
	}
	fmt.Printf("5 bytes: %s\n", string(b4))


	
	f.Close()
}