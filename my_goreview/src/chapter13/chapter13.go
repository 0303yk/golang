//this chapter telled 可变参数函数。sorry for my poor english
//谁敢说我的代码不能跑？不能跑我跑！
package main

import "fmt"

func check(a int , b ...int ){
	for _,v := range b {
		if v==a {
			fmt.Println("true",v)
		}else{
			fmt.Print(" ")
		}
	}
}

func main() {
	fmt.Println("第一个：")
	check(11,12,15,14,11)
	fmt.Println("第二个：")
	check(45, 56, 67, 45, 90, 109)//这行运行下来还有两个空格在最后，偷偷出现在后面。
}

//还是人家的专业啊！ 当然，他有弊端，在plus中表述
/*
package main

import (
    "fmt"
)

func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}
*/