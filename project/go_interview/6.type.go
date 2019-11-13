package main

import (
	"fmt"
)

type MyInt1 int 
type MyInt2 = int

func main() {
	var i int = 0
	var i1 MyInt1 = i 
	var i2 MyInt2 = i 

	fmt.Println(i1, i2)

	//cannot use i (type int) as type MyInt1 in assignment
	//类型别名和原类型是相同的，而类型定义和原类型是不同的两个类型。
	//不同类型的值不能赋值
}