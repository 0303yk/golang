package main
//与函数一起使用已定义类型
import "fmt"

type subscriber struct{
	name string
	rate float64
	active bool
}
//func printInfo 实现对subscriber 所有信息的输出，无返回值
func printInfo(s2 subscriber){
	fmt.Println("Name:",s2.name)
	fmt.Println("Monthly rate:",s2.rate)
	fmt.Println("Active?",s2.active)
}

//func defaultSubscriber 实现对subscriber的创建
func defaultSubscriber(name string)subscriber{
	var s1 subscriber
	s1.name = name
	s1.rate = 5.99
	s1.active = true
	return s1
}

func main(){
	subscriber1 := defaultSubscriber("Alan")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Bob")
	subscriber2.active = false
	printInfo(subscriber2)
}