package main

import (
	"fmt"
)

//下面是创建命名结构体
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
	fmt.Println(emp1)
	emp2 := imployee{"pines","dipper",29,1500 }
	fmt.Println(emp2)


	//声明匿名类
	empniming := struct {
		firstname string
		lastname string
		age int
		salary int
	}{
		firstname : "Mabel",
		lastname : "lover",
		age : 8,
		salary : 21000,
	}
	fmt.Println(empniming)
}
