//this package mainly talked about Select
package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	//s1 := <-output1
	//fmt.Println(s1)
	//s2 := <-output2
	//fmt.Println(s2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	//到这里可以自己思考这个东西的实际用途:哪个快用哪个?
	//yeah , i got the point !
	//官方解释:让我们假设我们有一个关键任务的应用程序，我们需要尽快将输出返回给用户。该应用程序的数据库被复制并存储在世界各地的不同服务器中。假设函数 server1 和 server2 实际上与 2 个这样的服务器通信。每个服务器的响应时间取决于每个服务器的负载和网络延迟。我们将请求发送到两个服务器，然后使用 select 语句在相应的 channel 上等待响应。首先响应的服务器由 select 选择，其他响应被忽略。这样我们就可以向多个服务器发送相同的请求，并将最快的响应返回给用户:)。
	//that's why function called server
}
