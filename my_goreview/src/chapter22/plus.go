//this packafe talked about the comparation
//容量是 channel 可以容纳的值的数量(make 语句中的capacity)
//长度是当前在其中排队的元素个数
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	//猜cap = 3 len = 2
	//BTW,参考上一节的死锁代码说明一个问题:
	//we can 在make里多创建capacity,但是创建少在下面设置通道时增多会出现deadlock

	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	//如下:每执行一个 len 就会减1 , 但是capacity不变
	fmt.Println("new length is", len(ch))
	//第三篇讲waitgroup(workpool的实现)
}
