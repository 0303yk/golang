package main

import "fmt"

func printbyte (s string){
	fmt.Printf("the byte is :")
	for i := 0 ; i<len(s) ; i++{

		fmt.Printf("%x ",s[i])
	}
}

func printunicode (s string){
	fmt.Printf("the char is :")
	for i := 0 ; i<len(s) ; i++{
		//这样的遍历和打印存在问题，即部分字符的UTF-8 encoding 编码占用两个字节。
		//此时讲rune
		fmt.Printf("%c ",s[i])
	}
}

func main() {
	name := "hello this is lily."
	printbyte(name)
	fmt.Printf("\n")
	printunicode(name)
}
