package main
//P456 action关键字 ,用于将数据插入模板
// 但是写入函数，range action 使用
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
	templatetext := "before loop: {{ .}}\n{{range.}}in loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executetemplate(templatetext,[]string{"do","re","mi"})
	templatetext2 := "{{range.}}in loop: {{.}}\n{{end}}\n"
	executetemplate(templatetext2,[]string{""})
	executetemplate(templatetext2,nil)

}
