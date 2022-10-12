package main
//P453的项目，使用Execute（写出模板文本），os.Stdout（将模板写入终端）
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
	text := "heres my template"
	tmpl,err := template.New("test").Parse(text)
	check(err)
	//Execute相当于println，会有输出的
	err = tmpl.Execute(os.Stdout,nil)
	check(err)

}
