package chapter13
//解决刚才即使优化过，仍旧不清楚对应关系的问题
//要点：channel后面可以添加基础类型，切片，映射，stuct复合类型，
//便于channel将两者一同发送，从而辨别清楚对应关系的问题
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	)

type Page struct {
	url string
	size int
}

func responsesize(url string,channel chan Page){
	//fmt.Println("getting:",url)
	r,err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer r.Body.Close()
	body , err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}
	channel <- Page{url:url , size:len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{
		"http://example.com/",
		"http://bd.com/",
		"https://tx.org/",
	}
	for _,url := range urls{
		go responsesize(url,pages)
	}
	for i := 0 ; i< len(urls);i++{
		page := <- pages
		fmt.Printf("%s: %d\n",page.url,page.size)
	}





}



