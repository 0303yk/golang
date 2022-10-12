package main


import ("fmt"
		"main4"
		)


func main(){
	var value myStruct

	value.Myfield = 3

	var pointer *MyStruct = &Value

	fmt.Println((*pointer).Myfield)

	// output: main.subscriber{name:"", rate:0, active:false}
	// fmt.Printf("%#v\n",s1)
}
