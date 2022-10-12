package main
//后续可能会用到的鬼东西，递归出目录内容
//Page  362

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error{
	fmt.Println(path)
	files,err := ioutil.ReadDir(path)
	if err != nil {
		//
		return err
	}
	for _,file := range files {
		filePath := filepath.Join(path,file.Name())
		if file.IsDir(){
			err := scanDirectory(filePath)
			if err != nil{
				//
				return err
			}
		}else{
			fmt.Println(filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("olded/src")
	if err != nil{
		log.Fatal(err)
	}
}
