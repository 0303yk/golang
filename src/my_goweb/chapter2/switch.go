//关键字：断言类型,不采用elif 用switch，case
package main

import (
	"fmt"
	"reflect"
	//"strconv"
)
type Element interface {}
type List [] Element

type Person struct {
	name string
	age int
}
//定义了 string方法,实现了fmt.stringer
func (p Person) string () string {
	return " (name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main () {
	list := make(List, 3)
	list [0] = 1       // an int
	list [1] = "Hello" // a string
	list [2] = Person{"Dennis", 70}
	for index, element := range list {
		//element.(type) 只能在switch中使用
		switch value := element.(type){
		//为啥下面一行代码不行？报错右面不能作为值传给value
		//switch value := reflect.TypeOf(element){
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			//挖个坑：把下面一行打印去除“main.”
			fmt.Printf("list[%d] is a Person and its value is %#v\n", index, value)
			//还有一个坑：为啥这里的fallthough无法使用
			//fallthrough
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
