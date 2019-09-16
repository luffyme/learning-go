package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//结构体
	//结构体里面装的是基础类型，切片，字典，数组以及其他类型的结构体
	//因为结构体的存在，Go语言的变量才有了更加丰富多彩的形式，Go语言程序的高楼大厦正是通过结构体一层层组装起来的。

	//1.结构体的定义
	//注意结构体内部变量大小写，首字母大写是公开变量，首字母是小写的为内部变量，只有属于同一个package的代码才能直接访问。
	type Circle struct {
		x int
		y int
		Radius int
	}

	//1.1 可以通过制定结构体内部字段的名称和初始值来初始化结构体，可以只制定部分字段的初值，甚至可以一个字段都不制定
	//那些没有制定初值的字段会自动化初始化为相应类型的零值
	var c Circle = Circle {
		x : 100,
		y : 100,
		Radius : 50, //注意这里的逗号不能少
	}
	fmt.Printf("%+v\n", c)

	//1.2 还可以不指定字段名称来顺序字段初始化，需要显示提供所有字段的初值，一个都不能少
	var c2 Circle = Circle{100, 100, 50}
	fmt.Printf("%+v\n", c2)

	//结构体变量和普通变量都有指针形式，使用取地址符就可以得到结构体的指针类型。
	var c3 *Circle = &Circle{100, 100, 50}
	fmt.Printf("%+v\n", c3)

	//1.3 还可以通过new函数来创建一个零值结构体，所有字段都被初始化为相应类型的零值
	var c4 *Circle = new(Circle)
	fmt.Printf("%+v\n", c4)

	//1.4 还可以通过下面形式创建
	var c5 Circle
	fmt.Printf("%+v\n", c5)

	//2.零值结构体和nil结构体
	//nil结构体是指结构体变量没有指向一个实际存在的内存，这样的指针变量只会占用1个指针的存储空间。
	//零值结构体会实实在在的占用内存空间，只不过每个字段都是零值，如果结构体里面字段非常多，那么这个内存空间占用肯定会会很大
	var c6 *Circle = nil
	fmt.Printf("%+v\n", c6)

	//3.结构体的内存大小
	//可以通过unsafe包提供了获取结构体内存占用的函数Sizeof。
	//Circle结构体在64位机器上占用了24个字节，每个int类型都是8个字节。
	var c7 Circle = Circle{Radius:50}
	fmt.Println(unsafe.Sizeof(c7))

	//4.结构体的拷贝
	//结构体之间可以相互赋值，本质上试一次浅拷贝，拷贝了结构体内部的所有字段。
	//结构体指针之间也可以相互赋值，在本质上也是一次浅拷贝，不过拷贝的仅仅指针地址值，结构体的内容是共享的。
	var c8 Circle = Circle{Radius:50}
	var c9 Circle = c8
	fmt.Printf("%+v\n", c8)
	fmt.Printf("%+v\n", c9)
	c8.Radius = 100
	fmt.Printf("%+v\n", c8)
	fmt.Printf("%+v\n", c9)

	var c10 *Circle = &Circle{Radius:50}
	var c11 *Circle = c10
	fmt.Printf("%+v\n", c10)
	fmt.Printf("%+v\n", c11)
	c10.Radius = 100
	fmt.Printf("%+v\n", c10)
	fmt.Printf("%+v\n", c11)

	//5.结构体的参数传递
	//函数调用时参数传递结构体变量，GO语言支持值传递，也支持指针传递。
	//值传递设计结构体字段的浅拷贝，指针传递会共享结构体内容，只会拷贝指针地址，规则上和赋值是等价的。

	//6.结构体方法
	//Go语言不是面向对象的语言，它里面不存在类的概念，结构体正是类的替代品，类可以附加很多成员方法，结构体也可以。
	//Go语言里面没有类型的隐形转换
	//Go语言里面没有self和this这样的关键字来指代当前的对象，他是用自己定义的变量名，通常都是使用单字母来表示。
	//GO语言的方法名称也分首字母大小写，首字母大写的就是公开方法，首字母小写就是内部方法，只能归属同一个包的代码才能访问。
	/*
		type Circle struct {
			x int
			y int
			Radius int
		}

		func (c Circle) Area() float64 {
			return math.Pi * float64(c.Radius) * float64(c.Radius)
		}

		func (c Circle) Circumference() float64 {
			return 2 * math.Pi * float64(c.Radius)
		} 

		func main() {
			var c = Circle{Radius:50}
			fmt.Println(c.Area(), c.Circumference())
		}
	*/

	//7.结构体指针方法
	//结构体指针方法和值方法在调用形式上没有区别，只不过一个可以改变结构体内部状态，另外一个不会。
	//通过指针访问内部的字段需要2次内存读取操作，第一步是取得指针地址，第二部是读取地址的内容。
	//在方法调用的时候，指针传递可以避免结构体的拷贝操作，结构体比较大的时候，这种性能差距比较明显。
	/*
		func (c *Circle) expand() float64 {
			c.Radius = 100
		} 
	*/

	//8.内嵌结构体
	//结构体作为一种变量可以放进另外一个结构体作为一个字段使用。这种内嵌结构体的形式在Go语言里面称之为组合。
	type Point1 struct {
		x int
		y int
	}
	type Circle1 struct {
		loc Point1
		Radius int 
	}
	var c12 = Circle1{
		loc : Point1 {
			x : 100,
			y : 100,
		},
		Radius : 50,
	}
	fmt.Printf("%+v\n", c12)
	fmt.Printf("%+v\n", c12.loc)
	fmt.Printf("%+v\n", c12.loc.x)

	//9.匿名内嵌结构体
	//内嵌结构体不提供名称，外面的结构体将直接继承内嵌结构体所有的内部方法和字段。
	//就好像把子结构体的一切全部揉进了父结构体一样。
	type Point2 struct {
		x int
		y int
	}
	type Circle2 struct {
		Point2 	//匿名内嵌结构体
		Radius int 
	}
	var c13 = Circle2{
		Point2 : Point2 {
			x : 100,
			y : 100,
		},
		Radius : 50,
	}
	fmt.Printf("%+v\n", c13)
	fmt.Printf("%+v\n", c13.x)
	fmt.Printf("%+v\n", c13.Point2.x)

	//10.Go语言的结构体没有多态性
	//多态是指父类定义的方法可以调用子类实现的方法，不同的子类有不同的实现，从而给父类的方法带来了多样的不同的行为。
	//Go语言不支持多态，外结构体的方法不能覆盖内部结构体的方法。
	/*
		type Fruit struct {}
		func (f Fruit) eat() {
			fmt.Println("eat fruit")
		}
		func (f Fruit) enjoy() {
			fmt.Println("smell first")
			f.eat()
			fmt.Println("clean finally")
		}	

		type Apple struct {
			Fruit
		}
		func (a Apple) eat() {
			fmt.Println("eat Apple")
		}

		func main() {
			var apple = Apple{}
			apple.enjoy()
		}
	*/
	
}