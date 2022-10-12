package chapter13
//
import "fmt"

func ji(j chan int){
	j <- 1
	j <- 3
}

func ou(o chan int){
	o <-2
	o <-4
}


func main() {
	o := make(chan int)
	j := make(chan int)
	go ji(j)
	go ou(o)
	fmt.Println(<- j)
	fmt.Println(<- o)
	fmt.Println(<- j)
	fmt.Println(<- o)


}
