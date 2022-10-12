package main
//P455 action关键字 ,用于将数据插入模板
// 但是写入函数，if action 使用
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
	tmpl,err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout,data)
	check(err)
}
func main(){
	//{{if .}}***** {{end} 如果data传入真值，**才可以输出，
	// 但false时{{if .}}***** {{end} 之外的仍可输出。
	executetemplate("start {{if .}}Dot is true! {{end}} finish\n",true)
	executetemplate("start {{if .}}Dot is true! {{end}} finish\n",false)

}
