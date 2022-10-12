package main
//net/http使用
import (
	"log"
	"net/http"
)

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}
//参数列表里的参数名放在函数体中 := 右边 ，而
// 返回值列表（允许多指返回）返回值名称放在 := 左边
// （和上一个不在同一行），然后return 或者使用 return 返回值
func viewhandler(w http.ResponseWriter, r *http.Request){
	message := []byte("signature list goes here!")
	_,err := w.Write(message)
	check(err)

}
func main(){
	http.HandleFunc("/guestbook",viewhandler)
	err := http.ListenAndServe("localhost:8080",nil)
	log.Fatal(err)
}
