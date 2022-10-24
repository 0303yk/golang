//this package tells slice's 动态长度,增加一个元素后长度+1，但是容量变为原来的二倍。
//顺带学一下slice的一些方法，len cap 和 append
package main

import (
	"fmt"
)
//所以下面创建了个啥？

//struct{
//	Length int
//	Capacity int
//	//指向第0个元素的指针的地址
//	ZerothElement *byte
//}

func f1 (numbers []int){
	for i := range numbers{
		numbers[i] -= 2
	}
}

func main() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	vegatable := []string{"potato","tomato","lobot"}//holy shit , i don't even konw how to spell 红萝卜
	fruit := []string{"watermalon","peach","group","orange"}
	//...运算符将一个切片添加到另一个切片上
	food := append(vegatable,fruit...)
	fmt.Println(food)
	//下面讲切片传递给函数（部分代码在main函数以外）
	nufirst := []int{45,12,13,14}
	f1(nufirst)
	fmt.Println(nufirst)
	//多维切片
	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

}

