package main
//升和毫升的转化
import "fmt"

type liters float64
type milliliters float64
type gallons float64
//该方法实现将毫升转化为升
func (i milliliters) Toliters()liters{
	return liters(i/1000)
}
//该方法实现将升转化为毫升
func (i liters) Tomilliliters()milliliters{
	return milliliters(i*1000)
}
func main(){
	l := milliliters(3)
	fmt.Println(l.Toliters())
	fmt.Printf("%0.1f 升是 %0.1f 毫升",liters(0.5),liters(0.5).Tomilliliters())

}