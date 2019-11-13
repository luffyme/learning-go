package main

import (
	"fmt"
)

func main() {
	//append1()
	append2()
}

/*
func append1() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)

	//不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。可以使用 make() 初始化之后再用。
	//同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
}
*/
func append2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

