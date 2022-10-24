//chapter19, well,actually. it's not chapter19,it's chapter21
//mainly talked about 并行和并发(go语言表现为并发(同一时间只能执行一件事))
//ok that's my job! it's that good? 大概记得go 进程不能放在最后一行
//关于channel进程喊话,下回分解
package main

import "fmt"

func pprint(a int){
	for i := a ; i >=0 ; i-- {
		fmt.Println("oh hello")
	}
}
func ppprint(a int){
	for i := a ; i >=0 ; i-- {
		fmt.Println("oh another hello")
	}
}

func main() {
	go ppprint(3)
	pprint(4)

}
