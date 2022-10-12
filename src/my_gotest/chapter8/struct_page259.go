package main
// 该go文件用来存储emp和sub的信息，address与前两个的属性不同，
//或者说address类型是前两者相同的信息，并打印。
import "fmt"

type employees struct{
	name string
	salary float64
	address
}
type subscriber struct{
	name string
	rate float64
	active bool
	address
}
type address struct{
	street string
	city string
	state string
	postalcode string
}

func main(){
	sub1 := subscriber{name:"Alan"}
	sub1.street = "123 Oak St"
	sub1.city = "Omaha"
	sub1.state = "NE"
	sub1.postalcode = "68111"
	fmt.Printf("%#v\n",sub1)

	emp1 := employees{name:"Joy"}
	emp1.street = "456 Elm St"
	emp1.city = "Postaland"
	emp1.state = "OR"
	emp1.postalcode ="97222"
	fmt.Printf("%#v\n",emp1)
	//fmt.Println("Street":,emp1.street)
}
