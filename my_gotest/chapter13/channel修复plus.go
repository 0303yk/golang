package chapter13
//代码优化
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
	urls := []string{"http://example.com/","http://bd.com/","https://tx.org/",}
	for _,url := range urls{
		go responhttp(url,res)
	}
	for i := 0 ; i<=len(urls);i++{
		fmt.Println(<- res)
	}





}
