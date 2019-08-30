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
}

