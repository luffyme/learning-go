package main

import (
	"fmt"
	"unsafe"
)

func main(){
	//字典

	//1.字典的创建
	//使用make函数创建的字典是空的，长度为0，内部没有任何元素。
	var m map[int]string = make(map[int]string)
	fmt.Println(m)

	//另外一种创建初始化元素的字典
	var m1 map[int]string = map[int]string{
		90: "优秀",
		80: "良好",
		70: "及格", 	//注意这里的逗号不可缺少
	}
	fmt.Println(m1)

	//字典变量支持类型推导
	var m2 = map[int]string{
		90: "优秀",
		80: "良好",
		70: "及格", 	//注意这里的逗号不可缺少
	}
	fmt.Println(m2)

	//如果你可以预知字典内部键值对的梳理，那么还可以给make函数传递一个整数值，通知运行时提前配好相应的内存，避免字典在长大过程中经历多次扩容
	var m3 = make(map[int]string, 16)
	fmt.Println(m3)

	//2.字典的读写
	//使用中括号来读写内部元素，使用delete函数删除元素
	var m4 = map[int]string{
		90: "优秀",
		80: "良好",
		70: "及格", 	//注意这里的逗号不可缺少
	}

	//读取元素
	var level string
	var ok bool

	level = m4[70]
	fmt.Println(level)

	//在读取的时候，如果key不存在，不会抛出异常，会返回value类型对应的零值。
	level = m4[50]
	fmt.Println(level)

	//可以使用字典提供的特殊语法,确认key是否真的存在
	level, ok = m4[50]
	if ok {
		fmt.Println(level)
	} else {
		fmt.Println("50 not exists")
	}

	//新增或者修改元素
	m4[60] = "刚刚及格"

	//删除元素
	delete(m4, 90)
	fmt.Println(m4)

	//3.字典的遍历
	//字典的遍历提供了两种方式，一种是需要携带value，一种是需要key。需要用到Go语言里面的range关键字。
	//Go语言字典里面是没有提供诸如keys()、values()这样的方法，如果你要获取key列表，需要自己循环
	var m5 = map[int]string{
		90: "优秀",
		80: "良好",
		70: "及格", 	//注意这里的逗号不可缺少
	}

	for score, name := range m5 {
		fmt.Println(score, name)
	}

	for score := range m5 {
		fmt.Println(score)
	}

	//4.字典不是线程安全的
	//Go语言里面内置的字典不是线程安全的，如果需要线程安全，需要使用锁来控制。

	//5.字典变量本质
	//字典变量里面存的是一个地址指针，指针指向字典的头部对象，字典变量占用的空间是一个指针的大小。
	//64位机器是8个字节，32位机器是4个字节。可以使用unsafe包里面的Sizeof函数来计算一个变量的大小
	fmt.Println(unsafe.Sizeof(m5))
	
}