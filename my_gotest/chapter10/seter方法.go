package main
//seter方法
import "fmt"
type Date struct{
	Year int
}
func (d *Date) SetYear(year int){
	d.Year = year
}
func main(){
	data := Date{ }
	data.SetYear(2020)
	fmt.Println(data.Year)

}