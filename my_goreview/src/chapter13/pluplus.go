//过几天请记得为啥是这样的输出
package main

import (
	"fmt"
)

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	//[Go world playground]
	fmt.Println(s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	//[Go world]
	fmt.Println(welcome)

}
