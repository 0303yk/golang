//that package contain much more unbelong to it!
//so this true package chapter25, let's talk about
// 互斥锁, 嗯 中文丝毫没有违和感.

package main

import (
	"fmt"
	"sync"
)
var x = 0

func increment(wg *sync.WaitGroup , mt *sync.Mutex) {
	//互斥锁是一种结构类型，我们在第 15 行创建了一个零值的 Mutex  变量 m。 在上面的程序中，我们修改了 increment 函数，使递增 x = x + 1 的代码位于 m.Lock() 和 m.Unlock() 之间。现在这段代码没有任何竞争条件，因为在任何时间点只允许一个Goroutine 执行这段代码。
	mt.Lock()
	x = x + 1
	mt.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0 ; i< 1000 ; i++{
		w.Add(1)

		go increment(&w,&m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
	
}

