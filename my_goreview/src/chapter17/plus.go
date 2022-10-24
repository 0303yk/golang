// this package talked about Value receivers in methods and Value arguments in functions
//方法比函数调用时多了一个地址相关调用
package main
//func  Area (a Type){func body}
//func (a Type)Area(){func body}
import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectanglein argument to area
	*/
	//area(p)

	p.area()//calling value receiver with a pointer
	//because it was looked like (*p).area()
}