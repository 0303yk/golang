package main

import (
	"fmt"
	"reflect"
)
func main(){
	var x float64 = 3.4
	//TypeOf returns the reflection Type that represents the dynamic type of i.
	//If i is a nil interface value, TypeOf returns nil.
	fmt.Println("type:",reflect.TypeOf(x))
	fmt.Println("type:",reflect.TypeOf(x).String())
}


