package main

import (
	"fmt"
)

func main() {
	//字符串

	//Go语言里面的字符串是字节串，英文字符占用1个字节，非英文字符占多个字节。
	//我们通常所说的字符通常指unicode字符，你可以认为所有的英文和汉字在unicode字符集中都有一个唯一的整数编号。
	//一个unicode通常用4个字节来表示，对应的Go语言中的字符rune占4个字节。
	
	//1.按字节遍历
	var s = "嘻哈china"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	//2.按字符rune遍历
	for codepoint, runevalue := range s {
		fmt.Printf("%d %d ", codepoint, int32(runevalue))
	}
	fmt.Println()

	//3.字符串是只读的
	//可以使用下标来读取字符串执行位置的字节，但是无法修改这个位置的字节内容。

	//4.切割
	//字符串在内存形式上比较接近切片，它可以像切一样进行切割获取子串，子串和母串共享底层字节数组。
	var s1 = "hello world"
	var s2 = s1[3:8]
	fmt.Println(s2)

	//5.字节切片和字符串的相互转换
	var s3 = "hello world"
	var s4 = []byte(s3)
	var s5 = string(s4)
	fmt.Println(s4)
	fmt.Println(s5)
}