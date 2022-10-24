package main

import (
	"fmt"
	"math/rand"
)
func main() {
	//多case实用
	letter := "a"
	switch letter{
	case "a","i","e","o","u":
		fmt.Println("it's a vowel")
	default:
		fmt.Println("it's not a vowel")
	}

	//使用label来break双循环的外层循环
	outloop:
		for i := 0 ;i < 3 ;i++{
			for j := 1 ;j < 4 ;j++ {
				fmt.Printf("i=%d , j=%d\n",i , j)
				if i==j {
					break outloop
				}
			}
		}
	//在最后一个case下禁止使用fallthough，他将导致编译错误
	randloop:
		for{
			switch i := rand.Intn(100);{
			case i %2 == 0 :
				fmt.Println("get  even number %d",i)
				break randloop
			}
		}


}
