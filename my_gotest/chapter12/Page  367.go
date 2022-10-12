package main
//Page  367
import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory(path string) {
	fmt.Println(path)
	files,err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _,file := range files {
		filePath := filepath.Join(path,file.Name())
		if file.IsDir(){
			scanDirectory(filePath)
		}else{
			fmt.Println(filePath)
		}
	}
}

func main() {
	scanDirectory("olded/src")
}
