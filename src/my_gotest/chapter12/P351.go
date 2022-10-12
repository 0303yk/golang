package main
//从文件中读取数字，重新访问

//下面是运行结果：
//Opening data.txt
//Closing file
//2022/10/02 16:15:41 strconv.ParseFloat: parsing "saerg": invalid syntax
//exit status 1

//解释：为啥把os释放资源改成函数？原因：改成函数方便调用deffer关键字（我猜的）
//关闭或者说释放资源的代码逻辑上应该在getfloat方法，或者函数whaterer，
//这样就必定导致如果get方法或函数在运行出错时无法关闭释放资源
//因此选择把close函数提前，并添加deffer让其必定运行。
import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strconv"
)

//OpenFlie 函数取代原来的os.Open调用
func OpenFile(fileName string)(*os.File,error){
	fmt.Println("Opening",fileName)
	return os.Open(fileName)
}
func CloseFile(file *os.File){
	fmt.Println("Closing file")
	file.Close()
}
// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
func main(){
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _,number := range numbers{
		sum +=number
	}
	sampleCount :=float64(len(numbers))
	fmt.Printf("Average: %0.2f\n",sum/sampleCount)
	fmt.Println((90.7+87.5+78.9)/3)
}


