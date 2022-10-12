package chapter15
//大迷糊的P433
//目标：写一个written函数，简化代码，同时包含多个web文件
//捋一捋其中的思路：
// write函数只是为了简化view和saullt里面重复的代码，view方法没有返回值（同sault函数下不赘述），
// http包内部的HandleFunc函数有两个参数：string类型的路径，和viewhandler函数作为一个参数（一级函数可做参数传递待会捋）
// viewhandler函数有两个参数：http包里的ResponseWriter类型接口（用于更新发送到浏览器的响应），
// 后面不知道什么玩意儿的指针(表示来自浏览器的请求的值)，
// main中ListenAndServe用于监听浏览器的请求并对其做出响应。

//好乱啊，再看一遍上一节的代码：HandleFunc接收路径（如"/hello"），viewhandler一级函数作为参数，
// 那viewhandler呢？main中ListenAndServe用于监听浏览器的请求并对其做出响应，
// 杨科现在放弃了，看后面的讲解吧，淦。


//func viewhandler(w http.ResponseWriter, r *http.Request){
//	message := []byte("hello,web!")
//	_,err := w.Write(message)
//	if err != nil{
//		log.Fatal(err)
//	}
//}
//func main(){
//	http.HandleFunc("/hello",viewhandler)
//	err := http.ListenAndServe("localhost:8080",nil)
//	log.Fatal(err)
//}

import (
	"net/http"
	"log"
)

func write(w http.ResponseWriter, m string){
	_,err := w.Write([]byte(m))
	if err != nil{
		log.Fatal(err)
	}
}

func viewhandler(w http.ResponseWriter,r *http.Request){
	write(w,"hello,web!")
}
func saulthandler(w http.ResponseWriter, r *http.Request){
	write(w,"sault,web!")
}
func main(){
	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	//	DefaultServeMux.HandleFunc(pattern, handler)
	//}
	http.HandleFunc("/hello",viewhandler)
	http.HandleFunc("/sault",saulthandler)
	err := http.ListenAndServe("localhost:8080",nil)
	log.Fatal(err)


}


