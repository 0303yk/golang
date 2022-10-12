package main
//nrt改进,html/template使用
import (
"log"
"html/template"
"net/http"
)

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}
func viewhandler(w http.ResponseWriter, r *http.Request){
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

