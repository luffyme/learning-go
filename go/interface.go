package main

import (
	"fmt"
)

func main() {
	//接口

	//Go语言的接口类型非常特别，其他语言需要在类的定义上显式实现了某些接口，才可以说这个类具备了接口的能力。
	//但是Go语言的接口是隐式的，只要结构体上定义的方法在形式上（名称、参数和返回值）和接口定义一样，那么这个结构体就自动实现了这个接口
	//我们就可以使用这个接口变量指向这个结构体对象。
	/*
		type Smellable interface {
			smell()
		}
		type Eatable interface {
			eat()
		}
		type Apple struct{}
		func (a Apple) smell() {
			fmt.Println("apple can small")
		}
		func (a Apple) eat() {
			fmt.Println("apple can eat")
		}

		func main() {
			var s1 Smellable
			var s2 Eatable
			var apple = Apple{}
			s1 = apple
			s1.smell()
			s2 = apple
			s2.eat()
		}
	*/

	//1.空接口
	//如果一个接口里面没有定义任何方法，那么它就是空接口，任意结构体都隐式的实现了空接口
	//Go语言为了避免用户重复定义很多空接口，它自己内置了一个，这个空接口的名称叫interface{}
	//空接口里面没有任何方法，所以它不具备任何能力，但是它可以容纳任意对象，它是一个万能容器。
	//比如一个字典的key是字符串， 但是希望value可以容纳任意类型的对象，这个时候就可以使用空接口类型interface{}
	var user = map[string]interface{}{
		"age" : 30,
		"address" : "shenzhen",
		"married" :  true,
	}
	fmt.Println(user)

	//需要类型转换
	var age = user["age"].(int)
	var address = user["address"].(string)
	var married = user["married"].(bool)
	fmt.Println(age, address, married)

	//2.接口变量的本质
	//在使用接口的时候，我们要将接口看成一个特殊的容器，这个容器只能容纳一个对象，只有实现了这个接口类型的对象才可以放进去。
	//接口变量作为变量也是需要占用内存空间的，
	//接口变量也是由结构体定义的，这个结构体包含两个指针字段，一个字段指向被容纳的对象内存，另一个字段指向一个特殊结构体itab。
	//这个特殊的结构体包含了接口类型信息和被容纳对象的数据类型信息。
	/*
		type iface struct {
			tab *itab  // 类型指针
			data unsafe.Pointer  // 数据指针
		}

		type itab struct {
			inter *interfacetype // 接口类型信息
			_type *_type // 数据类型信息
			...
		}
	*/

	//3.用接口来模拟多态
	//接口是一种特殊的容器，它可以容纳多种不同的对象，只要这些对象都同样实现了接口定义的方法。
	//如果我们将容纳的对象替换成另一个对象，那不就可以完成上一节我们没有完成的多态功能了么
	/*
		type Fruitable interface {
			eat()
		}

		type Fruit struct {
			Name string  // 属性变量
			Fruitable  // 匿名内嵌接口变量
		}

		func (f Fruit) want() {
			fmt.Printf("I like ")
			f.eat() // 外结构体会自动继承匿名内嵌变量的方法
		}

		type Apple struct {}

		func (a Apple) eat() {
			fmt.Println("eating apple")
		}

		type Banana struct {}

		func (b Banana) eat() {
			fmt.Println("eating banana")
		}

		func main() {
			var f1 = Fruit{"Apple", Apple{}}
			var f2 = Fruit{"Banana", Banana{}}
			f1.want()
			f2.want()
		}
	*/


	//4.接口的组合继承
	//接口的定义也支持组合继承，比如我们可以将两个接口定义合并为一个接口
	/*
		type Smellable interface {
			smell()
		}

		type Eatable interface {
			eat()
		}

		//这时 Fruitable 接口就自动包含了 smell() 和 eat() 两个方法
		type Fruitable interface {
			Smellable
			Eatable
		}	
	*/

	//5.接口变量的赋值
	type Rect struct {
		Width int
		Height int
	}
    var a interface {}
    var r = Rect{50, 50}
    a = r

    var rx = a.(Rect)
    r.Width = 100
    r.Height = 100
	fmt.Println(rx)
	
	//6.指向指针的接口变量
	//从输出结果中可以看出指针变量 rx 指向的内存和变量 r 的内存是同一份。
	//因为在类型转换的过程中只发生了指针变量的内存复制，而指针变量指向的内存是共享的。
	/*
		type Rect struct {
			Width int
			Height int
		}
		var a interface {}
		var r = Rect{50, 50}
		a = &r

		var rx = a.(Rect)
		r.Width = 100
		r.Height = 100
		fmt.Println(rx)
	*/
}	