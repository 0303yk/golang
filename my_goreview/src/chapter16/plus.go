//访问struct的各个字段
package main

import "fmt"

type imployee struct {
	firstname string
	lastname string
	age int
	salary int
}

func main() {
	emp1 := imployee{
		firstname : "Sam",
		age : 28,
		salary : 3000,
		lastname : "Anderson",
	}
	//shit 这个不能用（不会用）
	//for i := range emp1{
	//
	//}
	//输出说明是引用类型！
	emp1.salary = 2000
	fmt.Println(emp1.age)
	fmt.Println(emp1.salary)
	fmt.Println(emp1.lastname)
	fmt.Println(emp1.firstname)
	fmt.Println(emp1)

	emp2 := imployee{}
	//lastname 被赋予字符串的零值“”
	fmt.Println(emp2.lastname)
	//age被赋予int的 0
	fmt.Println(emp2.age)

}
