package main
//定义方法P275
//别忘了初衷，是因为要解决 函数名冲突而后增加对函数描述使代码变得繁琐 的问题
import "fmt"
//函数定义：func + 函数名 （ 参数名 参数类型 ）返回值类型{ }
//定义方法： func(接收器参数名 接收器参数类型)  方法名（）{ }

type mytype string

func (m mytype) sayHi(){
	fmt.Println("sayhi")
}
func main(){
	//如果mytype 没有传string类型值：
	// missing argument to conversion to mytype: mytype()
	//Page275中间内容对应这点
	value := mytype("a mytype value")
	value.sayHi()
}
