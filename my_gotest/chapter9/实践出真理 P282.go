package main
//实践出真理 P282，指针问题
//别忘了初衷，是因为要解决 函数名冲突而后增加对函数描述使代码变得繁琐 的问题
import "fmt"
//定义方法： func(接收器参数名 接收器参数类型)  方法名（）{ }

type Number int

func (n *Number) Display(){
	fmt.Println(*n)
}
func (n *Number) Double(){
	*n *= 2
}
func main(){
	number := Number(4)
	number.Double()//无需修改方法的调用
	number.Display()
}
