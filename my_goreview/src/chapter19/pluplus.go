//并 发or行 多个goroutine代码实现,现在猜是后者喽
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers()
	go alphabets()
	//哈哈,下面1000就无法全部输出了吧!!!
	// time.Sleep(1000 * time.Millisecond)
	// output: 1 a 2 3 b main terminated
	time.Sleep(3000 * time.Millisecond)
	//output: 1 a 2 3 b 4 c 5 d e main terminated
	fmt.Println("main terminated")
	//工作原理在同包中的一个图片?真是同时运行的!!!
	//图片说明:蓝色图像的第一部分代表 numbers Goroutine，
	// 栗色的第二部分代表 alphabets Goroutine，
	// 绿色的第三部分代表主程序 Goroutine，
	// 黑色的最后部分合并以上三个来告诉我们程序是如何工作的。
	// 每个框顶部的 0 ms，250 ms 等字符串表示以毫秒为单位的时间，
	// 输出在每个框的底部表示为 1,2,3 等等。蓝色框告诉我们在 250 ms 后打印 1，
	// 在 500 ms 后打印2，依此类推。最后一个黑框的底部的值为 1 a 2 3 b 4 c 5 d e main terminated，
	// 这也是该程序的输出。图像是我借花献佛的，希望你能够理解该程序的工作原理。
}
