package main

import "fmt"
//P433理解之道：P435
//也包含P436上，够消化好久了啊

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
