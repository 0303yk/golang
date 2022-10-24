//用字节切片构造字符串
package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	strtwo := string(runeSlice)
	fmt.Println(strtwo)



}