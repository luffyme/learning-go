package main

import (
	"fmt"
)

func main() {
	//切片是Go语言中最重要的数据结构，英文词汇叫slice。
	//切片变量包含三个域，分别是底层数组的指针，切片的长度length，切片的容量capacity
	//如果切片的长度和容量是相同的，则称切片是满容的。
	//切片变量是底层数组的视图。
	
	//1.切片的创建
	//使用make函数创建的切片内容是零值切片，也就是内部数组都是零值。
	//make函数创建切片，需要三个函数，分别是切片的类型，切片的长度和容量，第三个环境是可选的，如果不提供，则长度与容量相同。
	var s1 []int = make([]int, 5, 8)  	//长度为5，容量为8
	var s2 []int = make([]int, 8) 		//长度跟容量都是8
	s3 := make([]int, 8) 				//长度跟容量都是8
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	//还提供了另外一种创建切片的方法，允许赋初值。这种方式创建的切片是满容的。
	var s4 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s4, len(s4), cap(s4))

	//空切片
	//如果在创建切片的时候，容量和长度都是零，则该切片叫做空切片。
	var s5 []int
	var s6 []int = []int{}
	var s7 []int = make([]int, 0)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)

	//2.切片的赋值
	//切片的赋值是一次浅拷贝，拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容。
	var s8 = make([]int, 5, 8)
	for i := 0; i < len(s8); i++ {
		s8[i] = i + 1
	}
	var s9 = s8
	fmt.Println(s8)
	fmt.Println(s9)

	//修改新切片的值
	s9[0] = 255
	fmt.Println(s8)
	fmt.Println(s9)

	//3.切片的遍历
	//切片的遍历，在语法上和数组一样的，除了支持下标的形式，还支持range关键字
	var s10 = []int{1, 2, 3, 4, 5}
	fmt.Println("range遍历s10切片 通过index：")
	for index := range s10 {
		fmt.Println(index, s10[index])
	}
	fmt.Println("range遍历s10切片 通过index, value：")
	for index, value := range s10 {
		fmt.Println(index, value)
	}

	//4.切片的追加
	//切片是动态的数组，通过追加操作可以改变切片的长度，切片的每一次追加后都会形成新的切片变量。
	//如果底层数组没有扩容，那么追加前后的两个切片变量共享底层数组。如果底层数组扩容了，那么追加前后的底层数组是分离的不共享的。
	//如果底层数组是共享的，一个切片的变化会影响到另外一个切片，这点需要注意。
	var s11 = []int{1, 2, 3, 4, 5}
	fmt.Println("初始化s11切片变量：")
	fmt.Println(s11, len(s11), cap(s11))

	//对满容切片进行追加会分离底层数组
	var s12 = append(s11, 6)
	fmt.Println("s11追加6后的s11与s12切片：")
	fmt.Println(s11, len(s11), cap(s11))
	fmt.Println(s12, len(s12), cap(s12))

	//对非满容的切片进行追加会共享底层数组。
	var s13 = append(s12, 7)
	fmt.Println("s12追加7后的s12与s13切片：")
	fmt.Println(s12, len(s12), cap(s12))
	fmt.Println(s13, len(s13), cap(s13))

	//注意
	//对满容切片追加会导致数组发生扩容更换新的数组，但是旧数组并不会被立即销毁回收，因为老切片还指向旧数组。

	//5.切片的域是只读的。
	//我们刚才说切片的长度是可以变化的，那为什么又说切片是只读的？不矛盾吗？
	//这是为了提醒读者注意切片追加后形成了一个新的切片变量，而老的切片变量的三个域其实不会改变。改变的只是底层数组。
	//这里说的是切片的域是只读的，而不是说切片是只读的。
	//切片的域是组成切片变量的三个部分，分别是底层数组的指针，切片的长度，切片的容量。

	//6.切片的切割
	//切片的切割可以类比于字符串的子串，并不是要把切片割断，而是从母切片中拷贝出来一个子切片，子切片和母切片共享底层数组。
	var s14 = []int{1, 2, 3, 4, 5, 6, 7}
	//start_index, end_index, 不包含end_index
	var s15 = s14[2:5]
	fmt.Println("母切片与子切片：")
	fmt.Println(s14, len(s14), cap(s14))
	fmt.Println(s15, len(s15), cap(s15))

	//我们可以注意到s15切片的长度为3， 容量为5，既然切割前后共享底层数组，那为什么容量会不一样？
	//其实s15子切片的内部数据指针不再指向数组的开头了。子切片的容量大小是从中间位置知道切片末尾的长度，所以容量才变为7
	
	//子切片语法上要提供起始位置和结束位置。这两个位置都是可选的，
	//不提供起始位置，默认从母切片的初始位置开始（不是底层数组的初始位置），
	//不提供结束位置，默认结束到母切片的尾部。
	//切片不支持负数的位置
	var s16 = []int{1, 2, 3, 4, 5, 6, 7}
	var s17 = s16[:5]
	var s18 = s16[3:]
	var s19 = s16[:]
	fmt.Println("不同切割位置切割：")
	fmt.Println(s16, len(s16), cap(s16))
	fmt.Println(s17, len(s17), cap(s17))
	fmt.Println(s18, len(s18), cap(s18))
	fmt.Println(s19, len(s19), cap(s19))

	//对数组进行切割可以转换为切片，切片将原数组作为内部底层数组。
	var s20  = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s21 = s20[2:6]
	fmt.Println("对数组进行切割：")
	fmt.Println(s21, len(s21), cap(s21))
	s20[4] = 100
	fmt.Println("对数组进行修改后：")
	fmt.Println(s21, len(s21), cap(s21))

	//7.切片的深拷贝
	//Go 语言还内置了一个 copy 函数，用来进行切片的深拷贝。不过其实也没那么深，只是深到底层的数组而已。
	//如果数组里面装的是指针，比如 []*int 类型，那么指针指向的内容还是共享的。
	//copy 函数不会因为原切片和目标切片的长度问题而额外分配底层数组的内存
	//它只负责拷贝数组的内容，从原切片拷贝到目标切片
	//拷贝的量是原切片和目标切片长度的较小值 —— min(len(src), len(dst))，函数返回的是拷贝的实际长度
	var s22 = make([]int, 5, 8)
	for i:=0;i<len(s22);i++ {
		s22[i] = i+1
	}
	fmt.Println("copy前的原始切片：")
	fmt.Println(s22)
	var s23 = make([]int, 2, 6)
	var n = copy(s23, s22)
	fmt.Println("copy后的切片与长度：")
	fmt.Println(s23, n)

	//8.切片的扩容点
	//当比较短的切片扩容时，系统会多分配 100% 的空间，也就是说分配的数组容量是切片长度的2倍。
	//但切片长度超过1024时，扩容策略调整为多分配 25% 的空间，这是为了避免空间的过多浪费。
	var s24 = make([]int, 6)
	var s25 = make([]int, 1024)
	s24 = append(s24, 1)
	s25 = append(s25, 2)
	fmt.Println("扩容后的切片：")
	fmt.Println(len(s24), cap(s24))
	fmt.Println(len(s25), cap(s25))
}