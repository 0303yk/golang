package main

import ("fmt"
		"log"
	)

type Date struct{
	year int
	month int
	day int
}
//setter 方法
func (d *Date) Year()int {
	return d.year
}

func main(){
	data := Date{ }
	err := data.SetYear(2020) //为啥创建 err 等于data.SetYear捏？
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(data.Year()) //上面报错，此行代码不运行。

}