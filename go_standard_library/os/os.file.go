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


	//Seek 到一个文件中已知的位置并从这个位置开始进行读取。
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
	

	//没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = f.Seek(0, 0)
	if err != nil {
		fmt.Println("Seek file error : ", err)
	}


	//bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	if err != nil {
		fmt.Println("Peek file error : ", err)
	}
	fmt.Printf("5 bytes: %s\n", string(b4))

	//任务结束后要关闭这个文件（通常这个操作应该在 Open操作后立即使用 defer 来完成）。
	f.Close()

	//写文件
	d1 := []byte("hello\ngo\n")
	err = ioutil.WriteFile("/tmp/data1", d1, 0644)
	if err != nil {
		fmt.Println("ioutil.WriteFile error : ", err)
	}

	f2, err := os.Create("/tmp/data2")
	if err != nil {
		fmt.Println("os.Create error : ", err)
	}
	defer f2.Close()

	d2 := []byte{115, 111, 109, 101, 10}
    n2, err = f2.Write(d2)
	if err != nil {
		fmt.Println("Write error : ", err)
	}
	fmt.Printf("wrote %d bytes\n", n2)
	
	n3, err = f2.WriteString("writes\n")
	if err != nil {
		fmt.Println("Write error : ", err)
	}
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()

	w := bufio.NewWriter(f2)
	n4, err := w.WriteString("buffered\n")
	if err != nil {
		fmt.Println("WriteString error : ", err)
	}
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}