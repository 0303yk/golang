package main
//page 310 导出的方法像字段一样被提升
import "fmt"

type MyType struct{
	EmbeddedType
}
type EmbeddedType string
func (e EmbeddedType) ExportedMethod(){
	fmt.Println("111")
}
func (e EmbeddedType) unexportedMethod(){

}
func main() {
	fmt.Println("")
	
}
