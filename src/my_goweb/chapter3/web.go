package main

import (
	"fmt"
	"log"
	"net/http"
)

func gethttp(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintln(w,"hello man!")
}

func main(){
	http.HandleFunc("/hello",gethttp)
	err := http.ListenAndServe("localhost:8080",nil)
	if err!= nil {
		log.Fatal(err)
	}
}