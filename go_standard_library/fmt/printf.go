package main

import (
	"fmt"
	"os"
)

type point struct {
    x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p) 				//打印结构体，不带字段名，只有值
	fmt.Printf("%+v\n", p) 				//打印结构体，字段名，值
	fmt.Printf("%#v\n", p) 				//打印结构体，结构体值类型，字段名，值
	fmt.Printf("%T\n", p) 				//打印结构体值类型
	fmt.Printf("%t\n", true) 			//布尔类型
	fmt.Printf("%d\n", 123) 			//int类型
	fmt.Printf("%b\n", 14)				//int类型，二进制形式
	fmt.Printf("%c\n", 33) 				//int类型，对应的字符
	fmt.Printf("%x\n", 456) 			//int类型，十六进制编码
	fmt.Printf("%f\n", 78.9) 			//浮点数
	fmt.Printf("%e\n", 123400000.0) 	//科学计数法
	fmt.Printf("%E\n", 123400000.0)		//科学计数法
	fmt.Printf("%s\n", "\"string\"") 	//字符串
	fmt.Printf("%q\n", "\"string\"")    //不解析字符串里面的表达式
	fmt.Printf("%x\n", "hex this") 		//%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
	fmt.Printf("%p\n", &p) 				//指针的值
	fmt.Printf("|%6d|%6d|\n", 12, 345) 	//控制输出结果的宽度和精度，可以使用在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)	//指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)	//左对齐
	fmt.Printf("|%6s|%6s|\n", "foo", "b")		//控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")		//左对齐

	s := fmt.Sprintf("a %s", "string") 			//格式化字符串
	fmt.Println(s)
	
	fmt.Fprintf(os.Stderr, "an %s\n", "error")	//使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
}