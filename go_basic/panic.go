package main

import (
	"fmt"
)

func main() {
	//异常与捕捉
	//Go 语言提供了 panic 和 recover 全局函数让我们可以抛出异常、捕获异常。
	//panic抛出

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error catched", err)
		}
	}()
	panic("this is error")

	//有个值得注意的地方时，panic 抛出的对象未必是错误对象，而 recover() 返回的对象正是 panic 抛出来的对象，所以它也不一定是错误对象。
	/*
		func panic(v interface{})
		func recover() interface{}
	*/
	
}