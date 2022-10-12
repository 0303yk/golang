package main
//chapter 8    struct tpye
//section 1 创建函数类型的struct并调用
import (
	"fmt"
)
// define a car tpye
type part struct{
	description string
	count       int
}
func showTnfo (p part){
	fmt.Println("description:",p.description)
	fmt.Println("count:",p.count)
}

func main() {
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	//fmt.Println(showTnfo(bolts))
	showTnfo(bolts)
}
