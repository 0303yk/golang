//this package mainly talked about default
//解决
package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	output := make(chan string)

	go process(output)
//这是一个死循环啊!哦没事了,串行了!!!
//again 没错,它就是一个死循环
	for {
		time.Sleep(3000 * time.Millisecond)
		select{
		case v := <- output:
			fmt.Println("received value ",v)
			return
		default:
			fmt.Println("no value received")
		}
	}
//输出解释:
// go进程运行时,直接上time包 , 因此睡了10.5秒
// 因此每次都不会发生case事件, 在输出10次 no value received 后
// 开始打印字符串 received value 和其传递的值 v: process successful
// main函数return
//怎末感觉没怎么学过return关键字呢 是循环和if 那里学的吗?

//死锁问题,代码如下:
	//ch := make(chan string)
	//    select {
	//    case <-ch:
	//    }
//在上面的程序中，我们创建了一个 channel  ch。我们尝试从第 6 行中选择的这个 channel 读取。 select 语句将永远阻塞，因为没有其他 Goroutine 写入此 channel ，因此将导致死锁。该程序将在运行时发生报错，

//同时:如果存在 default case ，则不会发生此死锁，因为在没有其他情况准备就绪时将执行 default case 。上面的程序用下面的 default case 重写
//	ch := make(chan string)
//	select {
//	case <-ch:
//	default;
//	fmt.Println("default case executed")
//	}

//类似地，即使 select 只有零 channel ，也会执行 default case 。
	// var ch chan string
	//    select {
	//    case v := <-ch:
	//        fmt.Println("received value", v)
	//    default:
	//        fmt.Println("default case executed")
	//    }
//在上面的程序中，ch 是 nil，第 8 行我们试图从 ch 中的 select 中读取。如果default 情况不存在，则 select 将永远被阻塞并导致死锁。由于我们在 select 中有一个默认的情况，它将被执行并且程序将输出，

//当多个case同时准备好时, 其中一个将随即执行.

}
