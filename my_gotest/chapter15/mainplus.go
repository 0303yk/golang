package main

import "fmt"
//P436下，妈的不学了！
//声明sayhi函数
func sayHi(){
	fmt.Println("hi!")
}

func sayBye(){
	fmt.Println("Bye!")
}

func twice(a func()){
	a()
	a()
}

func main() {
	var funca func()
	funca = sayHi
	funca()
	twice(sayHi)
	twice(sayBye)
}
