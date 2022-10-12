package chapter13

//P391
//初衷：想通过并发提高代码运行速度，
// 但是由于不清楚可能运行的时间，而
// time.Sleep(time.Second)只能推测效果不好，
//同时go关键字不能接受一个返回值，所以使用channel。
//channel确保发送的gorountine在接收channel尝试使用该值前已经发送了该值
import "fmt"

func greeting(myChannel chan string){
	myChannel <- "hi"
}

func main() {
	my := make(chan string)
	go greeting(my)
	fmt.Println(<-my)//这里可以将my存储在变量中
}
//example:


