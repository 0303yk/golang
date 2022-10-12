package main
//P360递归
import "fmt"

func re(start int,end int) {
	fmt.Println(start)
	if start < end {
		re(start+1, end)
	}
}
func main() {
	re(1,3)
}
