package main
//P454 action关键字 ,用于将数据插入模板
import ("log"
	"os"
	"text/template"

)
func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}
func main(){
	text := "template start\n Action: {{.}}\ntemplate end\n"
	tmpl,err := template.New("test").Parse(text)
	check(err)
	//Execute相当于println，会有输出的
	err = tmpl.Execute(os.Stdout,"1123")
	check(err)
	err = tmpl.Execute(os.Stdout,"13")
	check(err)
	err = tmpl.Execute(os.Stdout,"23")
	check(err)

}
