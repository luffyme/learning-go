package main

import (
	"fmt"
	"errors"
)

func main() {
	//错误接口
	//Go 语言规定凡是实现了错误接口的对象都是错误对象，这个错误接口只定义了一个方法。
	//注意这个接口的名称，它是小写的，是内置的全局接口。
	//通常一个名字如果是小写字母开头，那么它在包外就是不可见的，不过 error 是内置的特殊名称，它是全局可见的。
	/*
		type error interface {
			Error() string
		}
	*/

	//编写一个错误对象很简单，写一个结构体，然后挂在 Error() 方法就可以了。
	/*
		type SomeError struct {
			Reason string
		}

		func (s SomeError) Error() string {
			return s.Reason
		}

		func main() {
			var err error = SomeError{"something happened"}
			fmt.Println(err)
		}	
	*/

	//Go 语言内置了一个通用错误类型，在 errors 包里。这个包还提供了一个 New() 函数让我们方便地创建一个通用错误。
	/*
		package errors

		func New(text string) error {
			return &errorString{text}
		}

		type errorString struct {
			s string
		}

		func (e *errorString) Error() string {
			return e.s
		}	
	*/
	var err = errors.New("something happened")
	fmt.Println(err.Error())

	//如果你的错误字符串需要定制一些参数，可使用 fmt 包提供了 Errorf 函数
	var thing = "something"
	var err2 = fmt.Errorf("%s happened 2", thing)
	fmt.Println(err2)
}