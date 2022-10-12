package chapter13
//P396
//这个文件解决的问题：刚才的检索web文件大小代码优化
//使用go并发，但不使用time，使其最快检索，完成输出
//运行完：确实好了点，但是可以优化，简化代码见下一个文件
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responhttp(url string,channel chan int){
	fmt.Println("getting:",url)
	r,err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer r.Body.Close()
	body , err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}
	channel <- len(body)
}

func main() {
	res := make(chan int)
	go responhttp("http://example.com/",res)
	go responhttp("http://baidu.com/",res)
	go responhttp("https://tx.org/",res)

	fmt.Println(<- res)
	fmt.Println(<- res)
	fmt.Println(<- res)



}
