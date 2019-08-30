package main

import (
	"fmt"
)

func main() {
	//3.for 循环
	//Go 语言虽然没有提供 while 和 do while 语句，不过这两个语句都可以使用 for 循环的形式来模拟。
	//平时使用 while 语句来写死循环 while(true) {}，Go 语言可以这么写
	/* 
	for {
        fmt.Println("hello world!")
	}
	*/
	//也可以这样写，效果是一样的
	/* 
	for true {
        fmt.Println("hello world!")
	}
	*/
	//for 什么条件也不带的，相当于 loop {} 语句。
	//for 带一个条件的相当于 while 语句
	//for 带三个语句的就是普通的 for 语句。
	for i := 0; i < 10; i++ {
		fmt.Println("hello world!")
	}
	
	//4.循环控制
	//Go 语言支持 continue 和 break 语句来控制循环，这两个语句和其它语言没什么特殊的区别。
	//除此之外 Go 语言还支持 goto 语句，我们作为普通用户平时还是尽量不要使用 goto 语句了。
}