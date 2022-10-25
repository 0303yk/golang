// chapter22 has a lot codes , so let's devide it at chapter23
//well maybe we found code about chapter 23 in 22,whatever.
//重写chapter19 的 plus 代码
package main

import (
	"fmt"
)

func hello(done chan bool ) {
	done <- true
	fmt.Println("hello function")

}

func main() {
	c := make(chan bool)
//创建了一个 done(AKA c) bool channel 。
//并将其作为参数传递给 hello Goroutine。我们正在从 done  channel 接收数据。
//这行代码是阻塞的，这意味着在 Goroutine 将数据写入 done  channel 之前，
//程序不会执行下一行代码。因此，我们就不需要原始程序中的 time.Sleep 了。
//代码行 <-done 从 done  channel 接收数据，但不在任何变量中使用或存储该数据。这是完全合法的。
//现在我们的 main Goroutine 被阻塞等待 done  channel 的数据。
// hello  Goroutine接收此 channel 作为参数，输出 Hello world  goroutine
// 然后写入 done  channel 。当这个写入完成时，main Goroutine 从 done  channel 接收数据，
// 它被解锁，然后输出 main function。
	fmt.Println("main function")
	go hello(c)
	<- c
	fmt.Println("main function")
}

//func hello() {
//	fmt.Println("Hello world goroutine")
//}
//
//func main() {
//	time.Sleep(5)
//	go hello()
//	fmt.Println("main function")
//}

