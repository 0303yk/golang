//yeah surprise m*f*!(go没有运行出来 why?)
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func main() {
	//holy shit ! guess what? i can use time package without doc's help
	time.Sleep(5)
	go hello()
	fmt.Println("main function")
}
//下一个包讲多个goroutines并发(maybe 这个词吧)
