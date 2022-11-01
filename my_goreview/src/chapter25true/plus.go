//this package talked about the use of channel 解决竞态条件
package main

import (
	"fmt"
	"sync"
)

var x = 0
//在上面的程序中，我们创建了一个容量为 1 的缓冲 channel ，并在 22 行中将其传递给increment Goroutine。此缓冲 channel 用于确保只有一个 Goroutine 访问增加 x 的代码的临界区。 这是通过在 x 递增之前将 true 传递给缓冲 channel 来完成的。 由于缓冲 channel 的容量为 1，因此在尝试写入此 channel 的所有其他 Goroutines 都会被阻塞，直到在递增 x 后从该 channel 读取值为止。这实际上只允许一个 Goroutine 访问临界区。
func increment(wg *sync.WaitGroup,ch chan bool){
	ch <- true
	x = x +1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool,1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
	
}
//summary ; 一般来说，当 Goroutine 需要彼此通信时使用 channel ，当只有一个 Goroutine 应该访问代码的临界区时使用互斥锁。