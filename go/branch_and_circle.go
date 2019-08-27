package main

import (
	"fmt"
)

func main() {
	//分支与循环

	//1.if else 语句
	//最大最小函数，我们一般会使用三元操作符 a>b?a:b 一条语句搞定。
	//不过 Go 语言没有三元操作符，这里只能使用 if 语句
	//我们还需要注意到 if else 语句的条件不需要括号扩起来
	var a = 21
    if a > 0 {
        fmt.Println("a大于0")
    } else if a < 0 {
        fmt.Println("a小于0")
    } else {
        fmt.Println("a等于0")
	}
	
	//2.switch 语句
	//switch 有两种匹配模式，一种是变量值匹配，一种是表达式匹配。
	//变量值匹配
	var score = 70
    switch score / 10 {
		case 0, 1, 2, 3, 4, 5:
			fmt.Println("成绩很差")
		case 6, 7:
			fmt.Println("成绩刚及格")
		case 8:
			fmt.Println("成绩很好")
		default:
			fmt.Println("成绩非常优秀")
	}
	//表达式匹配
    switch {
		case score < 60:
			fmt.Println("成绩很差")
		case score < 80:
			fmt.Println("成绩刚及格")
		case score < 90:
			fmt.Println("成绩很好")
		default:
			fmt.Println("成绩非常优秀")
	}

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

