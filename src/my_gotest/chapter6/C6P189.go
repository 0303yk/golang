package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
func main(){
	numbers, err := GetFloats("data.txt")
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

