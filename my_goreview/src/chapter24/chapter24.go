//well , this chapter actually is about waitgroup
package main
import (
	"fmt"            
	"sync"
	"time"
)
//仔细看下面函数的创建和使用 waitgroup
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	//创建了一个waitgroup后,
	//它等待所有 Goroutines 完成执行。
	//程序一直被阻塞直到所有 Goroutines 完成执行
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
