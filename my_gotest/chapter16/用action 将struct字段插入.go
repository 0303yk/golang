package main
//P456 action关键字 ,用于将数据插入模板
// 但是写入函数，使用action 将struct字段插入模板
import (
	"log"
	"os"
	"text/template"
)
type Subscriber struct{
	Name string
	Rate float64
	Active bool
}
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
	templatetext := "Name:{{.Name}}\n{{if.Active}}Rate:{{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name:"Alan",Rate:4.99,Active:true}
	executetemplate(templatetext,subscriber)//subscriber可以是任意类型，一级函数也可
	//上面赋值过了，所以下面只能是修改，不能再次赋值嗷
	subscriber = Subscriber{Name:"Bob",Rate:5.99,Active:false}
	executetemplate(templatetext,subscriber)
}

