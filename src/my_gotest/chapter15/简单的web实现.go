package chapter15
//P428
import ("net/http"
		"log")
//后续代码会给出viewhandler参数
// (w http.ResponseWriter, r *http.Request)
// 这样写的原因，但也请记得多复习前面的数据类型
func viewhandler(w http.ResponseWriter, r *http.Request){
	message := []byte("hello,web!")
	_,err := w.Write(message)
	if err != nil{
		log.Fatal(err)
	}
}
func main(){
	http.HandleFunc("/hello",viewhandler)
	err := http.ListenAndServe("localhost:8080",nil)
	log.Fatal(err)
}


