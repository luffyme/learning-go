package main

import (
	"fmt"
)

func main() {
	//变量

	//1.变量的定义
	//1.1变量最繁琐的定义
	//通过var关键字，显示定义变量，在变量名称s后面声明了变量的类型int
	var s int = 40
	fmt.Println(s)

	//1.2类型推导定义,将类型去掉，编译器会自动推导出变量类型
	var s1 = 41
	fmt.Println(s1)

	//1.3去掉var关键字。但是赋值的等号变成了:=,表示变量的自动类型推导+复制
	s2 := 42
	fmt.Println(s2)

	//2.全局变量与局部变量
	//如果变量定义在函数内部，函数调用结束变量就消亡了，与之对应的是全局变量，在程序运行期间，一直存在，它定义在函数外面。
	//如果全局变量的首字母大写，那么它就是公开的全局变量，如果全局变量的首字母是小写，那么它就是内部的全局变量，只能在当前包内的代码访问，外面的包是看不见的。

	//3.指针类型
	//指针符号 * 和取地址符号 &，同C语言一样，指针还支持二级指针，三级制针，只不过在日常应用中，很少遇到。
	//指针变量本质上就是一个整形变量，里面存储的值就是另一个变量内存的地址。
	var p int = 44
	var p1 *int = &p
	var p2 **int = &p1
	var p3 ***int = &p2
	fmt.Println(p1, p2, p3)
	fmt.Println(*p1, **p2, ***p3)

	//3.基础类型大全
	// 有符号整数，可以表示正负
	var a int8 = 1 // 1 字节
	var b int16 = 2 // 2 字节
	var c int32 = 3 // 4 字节
	var d int64 = 4 // 8 字节
	fmt.Println(a, b, c, d)

	// 无符号整数，只能表示非负数
	var ua uint8 = 1
	var ub uint16 = 2
	var uc uint32 = 3
	var ud uint64 = 4
	fmt.Println(ua, ub, uc, ud)

	// int 类型，在32位机器上占4个字节，在64位机器上占8个字节
	var e int = 5
	var ue uint = 5
	fmt.Println(e, ue)

	// bool 类型
	var f bool = true
	fmt.Println(f)

	// 字节类型
	var j byte = 'a'
	fmt.Println(j)

	// 字符串类型
	var g string = "abcdefg"
	fmt.Println(g)

	// 浮点数
	var h float32 = 3.14
	var i float64 = 3.141592653
	fmt.Println(h, i)
	
	//3.变量与常量
	//Go语言还提供了常量关键字const，用于定义常量，常量可以是全局常量，也可以是局部常量。
	//你不可以修改常量，否则编译器会报错。常量必须初始化。
	const c1 int = 43
	const c2 = 1000 
	const c3, c4 = 42, 1000

	//使用const()块定义
	const (					
		c5 = 42
		c6 = 1000
	)

	//如果常量未给出定义，则延续上面的定义规则
	const (					
		c7 = 42
		c8 // c8 = 42
		c9 // c9 = 42
	)

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9)

	//iota迭代定义常量
	//iota的规则是：若iota出现在const()中，那么const()定义的第一行的iota就是0，第二行就是0+1=1，不论iota是否被常量使用。
	const (
		gender_secret = iota
		gender_male // = iota
		gender_female // = iota
	)
	fmt.Println(gender_secret, gender_male, gender_female)

	const (
		c11 = 42 // iota = 0，虽然未使用iota，但后边（const()中）有使用，此时iota被初始化为0，下面每行累加1
		c12 = iota      // iota = 1，iota继续累加，使用了iota
		c13 = 1024      // iota = 2，同样未使用，但iota的值继续累加。c3 被定义为1024
		c14             // iota = 3，c4延续上面的定义c4=1024，iota继续累加
		c15 = iota      // iota = 4，iota继续累加，使用了iota
		c16 = iota      // iota = 5，iota继续累加，使用了iota
	)
	//此时结果为：42, 1, 1024, 1024, 4, 5
	fmt.Println(c11, c12, c13, c14, c15, c16)
	
}