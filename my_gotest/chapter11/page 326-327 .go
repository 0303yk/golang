package main
//page 326-327
import "fmt"

type Myinterface interface{
	MeWithoutPara()
	MeWithPara(float64)
	MeWithReturn()string
}
type MyType int

func (m MyType) MeWithoutPara(){
	fmt.Println("1")
}

func (m MyType) MeWithPara(f float64){
	fmt.Println("2",f)
}

func (m MyType) MeWithReturn()string{
	return "3"
}
func (my MyType) Mnot(){
	fmt.Println("4")
}
func main() {
	var value Myinterface
	value = MyType(5)
	value.MeWithoutPara()
	value.MeWithPara(127.3)
	fmt.Println(value.MeWithReturn())
}
