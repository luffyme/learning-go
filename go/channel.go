package main

import (
	"fmt"
)

func main() {
	//通道

	//1.创建通道
	//创建通道只有一种语法，那就是 make 全局函数，提供第一个类型参数限定通道可以容纳的数据类型，再提供第二个整数参数作为通道的容器大小。
	//大小参数是可选的，如果不填，那这个通道的容量为零，叫着「非缓冲型通道」，与之对应的有限定大小的通道就是缓冲型通道。
	//非缓冲型通道必须确保有协程正在尝试读取当前通道，否则写操作就会阻塞直到有其它协程来从通道中读东西。

	// 缓冲型通道，里面只能放整数
	//var bufferedChannel = make(chan int, 1024)
	// 非缓冲型通道
	//var unbufferedChannel = make(chan int)

	//2.读写通道
	//Go 语言为通道的读写设计了特殊的箭头语法糖 <-
	//把箭头写在通道变量的右边就是写通道，把箭头写在通道的左边就是读通道。
	//一次只能读写一个元素。

	var ch chan int = make(chan int, 4)
	for i:=0; i<cap(ch); i++ {
		ch <- i   // 写通道
	}
	for len(ch) > 0 {
		var value int = <- ch  // 读通道
		fmt.Println(value)
	}
	
	//2.读写阻塞
	//通道满了，写操作就会阻塞，协程就会进入休眠，直到有其它协程读通道挪出了空间，协程才会被唤醒。
	//如果有多个协程的写操作都阻塞了，一个读操作只会唤醒一个协程。

	//通道空了，读操作就会阻塞，协程也会进入睡眠，直到有其它协程写通道装进了数据才会被唤醒。
	//如果有多个协程的读操作阻塞了，一个写操作也只会唤醒一个协程。

	/*
		func send(ch chan int) {
			for {
				var value = rand.Intn(100)
				ch <- value
				fmt.Printf("send %d\n", value)
			}
		}

		func recv(ch chan int) {
			for {
				value := <- ch
				fmt.Printf("recv %d\n", value)
				time.Sleep(time.Second)
			}
		}

		func main() {
			var ch = make(chan int, 1)
			// 子协程循环读
			go recv(ch)
			// 主协程循环写
			send(ch)
		}
	*/

	//3.关闭通道
	//Go 语言的通道有点像文件，不但支持读写操作， 还支持关闭。
	//读取一个已经关闭的通道会立即返回通道类型的「零值」，而写一个已经关闭的通道会抛异常。
	//如果通道里的元素是整型的，读操作是不能通过返回值来确定通道是否关闭的。

	//可以使用 for range 语法糖来遍历通道
	//当通道空了，循环会暂停阻塞，当通道关闭时，阻塞停止，循环也跟着结束了。当循环结束时，我们就知道通道已经关闭了。
	var chnew = make(chan int, 4)
	chnew <- 1
	chnew <- 2
	close(chnew)

	// for range 遍历通道
	for value := range chnew {
		fmt.Println(value)
	}
	
	//4.通道写安全
	//上面提到向一个已经关闭的通道执行写操作会抛出异常，这意味着我们在写通道时一定要确保通道没有被关闭。
	//Go 语言并不存在一个内置函数可以判断出通道是否已经被关闭。
	//即使存在这样一个函数，当你判断时通道没有关闭，并不意味着当你往通道里写数据时它就一定没有被关闭，并发环境下，它是可能被其它协程随时关闭的。

	//确保通道写安全的最好方式是由负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。

	//这个方法确实可以解决单写多读的场景，可要是遇上了多写单读的场合该怎么办呢？
	//任意一个读写通道的协程都不可以随意关闭通道，否则会导致其它写通道协程抛出异常。
	//这时候就必须让其它不相干的协程来干这件事，这个协程需要等待所有的写通道协程都结束运行后才能关闭通道。
	//那其它协程要如何才能知道所有的写通道已经结束运行了呢？这个就需要使用到内置 sync 包提供的 WaitGroup 对象，它使用计数来等待指定事件完成。

	//5.多路通道
	//在真实的世界中，还有一种消息传递场景，那就是消费者有多个消费来源，只要有一个来源生产了数据，消费者就可以读这个数据进行消费。
	/*
		func send(ch chan int, gap time.Duration) {
			i := 0
			for {
				i++
				ch <- i
				time.Sleep(gap)
			}
		}

		func recv(ch1 chan int, ch2 chan int) {
			for {
				select {
					case v := <- ch1:
						fmt.Printf("recv %d from ch1\n", v)
					case v := <- ch2:
						fmt.Printf("recv %d from ch2\n", v)
				}
			}
		}

		func main() {
			var ch1 = make(chan int)
			var ch2 = make(chan int)
			go send(ch1, time.Second)
			go send(ch2, 2 * time.Second)
			recv(ch1, ch2)
		}	
	*/

	//6.非阻塞读写
	//Go 语言还提供了通道的非阻塞读写。当通道空时，读操作不会阻塞，当通道满时，写操作也不会阻塞。非阻塞读写需要依靠 select 语句的 default 分支。
	//当 select 语句所有通道都不可读写时，如果定义了 default 分支，那就会执行 default 分支逻辑，这样就起到了不阻塞的效果。
	/*
		for {
			select {
				case ch1 <- i:
					fmt.Printf("send ch1 %d\n", i)
				case ch2 <- i:
					fmt.Printf("send ch2 %d\n", i)
				default:
			}
		}
	*/
}