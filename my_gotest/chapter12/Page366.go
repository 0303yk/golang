package main
//Page366 延迟调用在panic前完成
import "fmt"

func main() {
	one()
}
func one(){
	defer fmt.Println("deffer belong one")//2
	two()
}
func two(){
	defer fmt.Println("deffer belong two")//1
	panic("oh no,,,,")//3
}