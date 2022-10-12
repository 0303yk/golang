package chapter13
//它好智能，甚至爬取了一模一样的网页
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r,err := http.Get("https://baidu.com/")
	if err != nil{
		log.Fatal(err)
	}
	defer r.Body.Close()
	body , err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
