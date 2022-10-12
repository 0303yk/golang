package main
import "fmt"
type Foo struct {name string}
// Set方法
func (f *Foo) SetName(name string) {
	f.name = name
}
// Get方法，无需使用Get，只要把首字母大写即可。
func (f Foo) Name() string {
	return f.name
}
func main() {
	p := Foo{}
	p.SetName("Abc")
	name := p.Name()
	fmt.Println(name)
	fmt.Println(p)
}
