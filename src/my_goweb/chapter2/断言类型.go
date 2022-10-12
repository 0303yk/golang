package main
//关键字：断言类型
import (
	"fmt"
	"strconv"
	//"container/heap"
	//"io"
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
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
