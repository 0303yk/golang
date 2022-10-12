package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("D:/go_workplace/olded/src/my_goweb/chapter4/login.gtpl")
		t.Execute(w,nil)
	} else{
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)

	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme:",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value",strings.Join(v," "))
	}
	fmt.Fprintln(w,"hello astaxie!")
}

func main(){
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login)
	err := http.ListenAndServe("localhost:9090",nil)
	if err!= nil {
		log.Fatal(err)
	}
}