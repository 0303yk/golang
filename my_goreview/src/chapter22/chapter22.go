//this package mainly talked about Channels
//create a channel with a buffe

//channel 可以看作是管道，Goroutines 使用
//管道进行通信。和水在管道中从一端流向另一端一样
//，数据可以通过管道从一端发送，从另一端接收。

package main
//下面简单的例子加了注释,另一个例子在下面.
//import "fmt"
//
//func main() {
//	//ch := make(chan type, capacity)
//	ch := make(chan string,2)
//	ch <- "naveen"
//	ch <- "paul"
//
//	fmt.Println(<- ch)
//	fmt.Println(<- ch)
//}

//another example

import (
	"fmt"
	"time"
)

func write(ch chan int){
	for i := 0 ;i < 5 ; i++{
		ch <- i
		fmt.Println("successfully wrote", i , "to ch")
	}
	close(ch)
	//fatal error: all goroutines are asleep - deadlock!
	//goroutine 1 [chan receive]:
	//main.main()
}

func main(){
	ch := make(chan int,2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch{
		fmt.Println("read value",v,"from ch")
		time.Sleep(2 * time.Second)
	}
	//孔雀东南飞,两秒一回头!
	//successfully wrote 0 to ch
	//successfully wrote 1 to ch

	//read value 0 from cha
	//successfully wrote 2 to ch

	//read value 1 from cha
	//successfully wrote 3 to ch

	//read value 2 from cha
	//successfully wrote 4 to ch

	//read value 3 from cha
	//read value 4 from cha

	//deadlock :
	//当将 3 个字符串写入容量为 2 的缓冲 channel 。当代码执行到第 11 行时，由于 channel 已超过其容量，因此写入被阻塞。
	// 现在 Goroutine 必须从 channel 中读取才能继续写入，但在这种情况下，没有 goroutine 从该 channel 同时读取。
	// 因此会出现 deadlock（死锁),代码example:

	//func main() {
	//    ch := make(chan string, 2)
	//    ch <- "naveen"
	//    ch <- "paul"
	//    ch <- "steve"
	//    fmt.Println(<-ch)
	//    fmt.Println(<-ch)
	//}
}


