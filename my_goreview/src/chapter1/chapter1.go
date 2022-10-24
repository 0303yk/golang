//this go package introduce the basic function about go. ok,i don't know what i said.
package main

import (
	"fmt"
	"math"
	"unsafe"
	)
func main() {
	var age int = 20
	var width,height int = 100, 50
	var (
		names = "Dipper"
		ages = "23"
		heights int
	)
	count := 10
	a, b := 4.0,5.0
	c := math.Min(a,b)

	testone := true
	testtwo := false
	judgefirst := testone&&testtwo
	judgesecond := testone||testtwo
	var unsafeone int = 89
	unsafetwo := 95
	complexone := 2+2i
	complestwo := complex(2,2)
	calcomp := complexone * complestwo
	//字符串拼接
	firstname := "Dave"
	lastname := "Kepofier"
	guyname := firstname +" "+lastname
	//这里开始，讲常量,第一点，常量不可以被二次赋值,布尔常量也是无类型常量
	const hello = "Hello World"
	var notypename = "Sam" //poor Sam
	//创建一个类型化常量
	const typedhello string = "Hello World again！"




 	println("hello world")
	fmt.Println("xiaoming is ", age)
	fmt.Println("width is ",width,"and height is " ,height)
	fmt.Println("name later is ",names,"and new age is", ages,"heights is ",heights)
	fmt.Println("count is",count)
	fmt.Println("the minimum value is:",c)
	fmt.Println("testone:",testone,"testtwo:",testtwo)
	fmt.Println("judgefirst",judgefirst)
	fmt.Println("judgesecond",judgesecond)
	fmt.Println("value of unsafeone is",unsafeone,"and unsafetwo is",unsafetwo)
	fmt.Printf("type of one is %T, size of one is %d",unsafeone,unsafe.Sizeof(unsafeone))
	//32位操作系统中，两者各占用32位（4字节）
	fmt.Printf("\ntype of two is %T, size of two is %d",unsafetwo,unsafe.Sizeof(unsafetwo))
	fmt.Println("\n复数的计算结果是：",calcomp)
	fmt.Println("the guyname's name is",guyname)
	//类型转换,当然在赋值操作也可使用强制类型转换
	weirdnum := age+int(a)
	fmt.Println("weirdnum man's new age is:",weirdnum)
	//这里开始，讲常量
	fmt.Println("it must be a const:",hello)
	fmt.Printf("type is %T, value is %v\n",notypename,notypename)
	fmt.Printf("new type is %T, value is %v\n",typedhello,typedhello)

	fmt.Println()
	fmt.Println()

	//下面这行代码，所以为啥math.Sqrt是汇编实现的？为啥不是我想的那种方法实现呢？
	fmt.Println(math.Sqrt(4))


}

