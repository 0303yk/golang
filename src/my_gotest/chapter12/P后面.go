package main

import "fmt"

func calmdown(){
	recover()
}

func breakout(){
	defer calmdown()
	panic("oh no")
}

func main() {
	breakout()
	fmt.Println("ewg")
	fmt.Println(recover())
}
