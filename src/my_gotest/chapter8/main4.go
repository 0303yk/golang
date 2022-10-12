package src
//使用函数修改structP242

import "fmt"

type subscriber struct{
	name string
	rate float64
	active bool
}
//func applyDiscount 函数用来设置折扣
func applyDiscount(s *subscriber){

	//4.99字段是struct中rate的拷贝，因此在main函数中输出原struct
	s.rate = 4.99
}

func main(){
	var s1 subscriber
	applyDiscount(&s1)

	fmt.Println(s1.rate)

	// output: main.subscriber{name:"", rate:0, active:false}
	// fmt.Printf("%#v\n",s1)
}
