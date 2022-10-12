//struct 字段创建并调用，算是复习了吧
//method 创建方法，避免函数名冗余、
//interface接口，是一组方法的集合
//断言 可以对接口类型进行判断，（双反回值下）
// 如果true，那么就可以调用其方法，无需知道其类型
package main

import (
	"fmt"
	"math"
)

type Human struct{
	Name string
	Id float64
}

type Teacher struct {
	people Human
	student Student
}

type Student struct{
	people Human
	score int
}

type sanjiao struct{
	di float64
	high float64
}

type juxing struct{
	di float64
	high float64
}

type yuan struct{
	banjing float64

}

func (s sanjiao) area() float64{
	return s.di*s.high
}

func (s juxing) area()float64{
	return s.di*s.high
}

func (s yuan) area()float64{
	return s.banjing*s.banjing*math.Pi
}

func main(){
	c1 := juxing{5,8}
	c2 := yuan{3}
	//stu1 := Student{Human{"Nike",410402},80}
	//fmt.Println(stu1.score)
	//fmt.Println(stu1.people.Name)
	fmt.Println(c1.area())
	fmt.Println(c2.area())

}
