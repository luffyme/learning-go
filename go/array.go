package main

import (
	"fmt"
)

func main() {
	//Go语言里面的数组其实不常用，因为数组是定长的静态的，一旦定义好长度就无法更改，用起来很不方便。

	//1.变量的定义
	//如果只声明不赋值，编译器会给数组默认赋上[零值]，数组的零值就是内部元素的零值
	var a [9]int

	//另外三种变量定义的方式，效果是一样的
	var b = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var c [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(a, b, c, d)

	//2.数组的访问
	//len可以获取数组的长度，数组长度是在编译期间确定的。
	var squares [9]int
	for i := 0; i < len(squares); i++ {
		squares[i] = (i + 1) * (i + 1)
	}

	fmt.Println(squares)

	//3.数组赋值
	//同样子元素类型并且是相同长度的数组才可以相互赋值，否则就是不同的数组类型，不能赋值。
	//数组的赋值本质上是一种浅拷贝操作， 赋值的两个数组变量的值不会共享。
	var e = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var f [9]int
	f = e 
	e[0] = 12345
	fmt.Println(e)
	fmt.Println(f)
	
	//4.数组遍历
	//数组除了可以使用for语句使用下标遍历外，还可以使用range关键字来遍历。
	var g = [5]int{1, 2, 3, 4, 5}
	for index := range g {
		fmt.Println(index)
	}
	for index, value := range g {
		fmt.Println(index, value)
	}
}