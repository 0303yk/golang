package main
//chapter 8    struct tpye
//section 2 编写minimumOrder函数，实现功能如代码21行所示：存储为一个新的struct
import "fmt"

type part struct{
	description string
	count       int
}

func minimumOrder(description string) part {
	var p part
	p.description=description
	p.count = 100
	return p
}
func main() {
	p := minimumOrder("Hex bolts")
	fmt.Println(p.description,p.count)
	// fmt.Printf("%#v\n",p)
	// main.part{description:"Hex bolts", count:100}
}

