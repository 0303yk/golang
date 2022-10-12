package main
//P455 action关键字 ,用于将数据插入模板
// 但是写入函数，劳资最烦创建函数了，哈哈，学吧.
import (
	"log"
	"os"
	"text/template"
)
func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}
func executetemplate(text string, data interface{ } ){
	//text := "template start\n Action: {{.}}\ntemplate end\n"报错出现在这一行，为啥我把它搞进来了？
	tmpl,err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout,data)
	//在模板action中使用给定的数据值
	//still：Execute相当于println，会有输出的
	check(err)
}
func main(){
	//interface类型的data保证可以传多种类型的值
	executetemplate("hei man\ndata is:{{.}}\n","123")
	executetemplate("hei man\ndata is:{{.}}",456)

}
