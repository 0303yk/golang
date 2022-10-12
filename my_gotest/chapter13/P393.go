package chapter13

import "fmt"

func greeting(myChannel chan string){
	myChannel <- "hi"
}

func main() {
	my := make(chan string)
	//my <- "hihi"
	go greeting(my)
	//greeting(my)
	fmt.Println(<-my)//这里可以将my存储在变量中
}
