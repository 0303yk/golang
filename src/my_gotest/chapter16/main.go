package main
//网址：http://localhost:8080/guestbook
//P460 保存签名myself
import(
	"os"
	//"fmt"
	"html/template"
	"net/http"
	"log"
	"bufio"
)
//func check用于检查err并返回

type Guestbook struct {
	SignatureCout int
	Signatures []string
}

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}
func viewhandler(w http.ResponseWriter, r *http.Request){
	signatures := getStrings("signature.txt")
	html,err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCout: len(signatures),
		Signatures: signatures,
	}
	err = html.Execute(w,guestbook)
	check(err)
}
func main(){
	http.HandleFunc("/guestbook",viewhandler)
	err := http.ListenAndServe("localhost:8080",nil)
	log.Fatal(err)
}