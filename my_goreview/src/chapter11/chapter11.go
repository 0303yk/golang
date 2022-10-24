//this package tell the different betwen arr and slice
package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("  ")
	}
}

func main() {
	var f [3]int //creat a arr with length 3
	var e [1] int
	e[0] = 3
	c :=  [4]float64{1.0,25.1,14.1,44.3}
	d :=  [...]float64{11.30,25.1,14.1,44.3} // creat arr which let itself found len
	//and 数组的大小是类型的一部分，数组是值类型
	//new
		a := [3][2]string{
			{"lion", "tiger"},
			{"cat", "dog"},
			{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
		}
		printarray(a)
		var b [3][2]string
		b[0][0] = "apple"
		b[0][1] = "samsung"
		b[1][0] = "microsoft"
		b[1][1] = "google"
		b[2][0] = "AT&T"
		b[2][1] = "T-Mobile"
		fmt.Printf("\n")
		printarray(b)

	fmt.Println(f)
	fmt.Println(e)
	fmt.Println(c)
	fmt.Println(d)
}
