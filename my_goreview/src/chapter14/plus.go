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
	runes := []rune(s)
	for i := 0 ; i<len(runes) ; i++{
		fmt.Printf("%c ",runes[i])
	}
}

func printCharAndBytes(s string){
	fmt.Println()
	for index , runes := range s {
		fmt.Printf("%c start at byte %d  ",runes,index)
	}
}

func main() {
	name := "hello this is lily."
	printbyte(name)
	fmt.Printf("\n")
	printunicode(name)
	nametwo := "Señor"
	//ñ 占用两个byte
	printCharAndBytes(nametwo)
}
