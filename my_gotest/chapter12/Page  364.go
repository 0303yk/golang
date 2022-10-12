package main
//Page  364
//发起一个panic:动机：解决在err判断中递归调用scanDirectory函数
// 时若出现无权限等问题报错，后续扫描均出错

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
