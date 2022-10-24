//this package mainly talk about interface
//大概这样吧
/*
func Area(a, b int)int{result := a*b
	return result}
type myinter interface{Area()}
 */

//好了，现在第十八章断片了，下面的吧，接口现在碰不到。give up！
package main

import (
	"fmt"
)

type VowelsFinder interface{
	FindVowels() []rune
}

type Mytstring string

func (ms Mytstring)FindVowels()[]rune {
	var result []rune
	for _,rune := range ms {
		if rune == 'a'|| rune == 'e'||rune == 'i'|| rune =='o'||rune == 'u'{
			result = append(result,rune)
		}
	}
return result
}

func main() {
	name := Mytstring("Sam Nike")
	var v VowelsFinder
	v = name
	fmt.Printf("%c",v.FindVowels())
}
