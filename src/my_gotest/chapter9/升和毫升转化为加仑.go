package main
//升和毫升转化为加仑
import "fmt"

type liters float64
type milliliters float64
type gallons float64
//该方法实现将升转化为加仑
func (i liters) Togallon()gallons{
	return gallons(i*0.264) // 在方法体中声明的 i 可以在花括号中使用？这样？ 所以第九行说明返回gallons类型，为什么可以使用？ gallons(i*0.264) 意思是强转liters为gallons，get it！
}
//该方法实现将毫升转化为加仑
func (j milliliters) Togallon()gallons{
	return gallons(j*0.000264)//同第10行理，liters(3.0)也可以转float为liters类型，ok
}
func main(){
	soda := liters(1.5)
	fmt.Printf("%0.3f 升苏打是 %0.3f 加仑\n",soda,soda.Togallon())
	water :=milliliters(500)
	fmt.Printf("%0.3f 升水是 %0.3f 加仑",water,water.Togallon())

}