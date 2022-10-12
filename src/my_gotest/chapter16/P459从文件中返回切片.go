package main
//now,got new keyboard，总是低头可不行啊，所以这节讲（P459之前所有整合）了什么？
//
import(
	"os"
	"fmt"
	"html/template"
	"net/http"
	"log"
	"bufio"
)
//func check用于检查err并返回

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func getstring(filename string )[]string {
	var lines []string
	file,err := os.Open(filename)
	if os.IsNotExist(err){
		return nil
	}
	check(err)
	defer file.Close()
	scanner :=bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines,scanner.Text())
	}
	check(scanner.Err())
	return lines
}
func viewhandler(w http.ResponseWriter, r *http.Request){
	signatures := getstring("signature.txt")
	fmt.Printf("%#v\n",signatures)
	html,err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(w,nil)
	check(err)
}
func main(){
	http.HandleFunc("/guestbook",viewhandler)
	err := http.ListenAndServe("localhost:8080",nil)
	log.Fatal(err)
}