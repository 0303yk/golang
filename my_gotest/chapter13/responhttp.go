package chapter13
//
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responhttp(url string){
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
	fmt.Println(len(body))
}

func main() {
	responhttp("http://example.com/")
	responhttp("http://baidu.com/")
	responhttp("https://tx.org/")



}
