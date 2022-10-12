package main
//p332pointer
import "fmt"

type Switch string

func (s *Switch) toggle(){
	if *s == "on"{
		*s = "off"
	}else {
		*s = "on"
	}
	fmt.Println(*s)
}
type Toggleadle interface{
	toggle()
}

func main() {
	s := Switch("on")
	var t Toggleadle = &s
	t.toggle()
	t.toggle()
}
