package main


//如果导入的包名前加_ ，则表示可以不调用包中的函数，但是该包中的init函数会被调用。
import (
	//_"chapter6/rectangle"
	"fmt"
)

func add(a int, b int)(int,bool){
	return a+b,a==b
}

func main() {
	var a = 6
	b := 7
	c,d := add(a,b)
	print(c,"  ",d,"\n")
	//if num := 10 ; num%2 == 0 {
	//	fmt.Println("even")
	//// else 必须与上一个if 或者 else if 语句之后的 } 的同一行开始
	//}else{
	//	fmt.Println("odd")
	//}
	//上面的写法不好，go哲学中要：
	num := 10
	if num %2 == 0{
		fmt.Println("even")
		return
	}
	fmt.Println("odd")




}
