package main
//使用函数进行类型转换
import ("fmt"
)


type jialun float64
type sheng float64

func tojialun(i sheng)jialun{
	return jialun(i*0.264)
}

func tosheng(j jialun)sheng{
	return sheng(j*3.785)
}

func main(){
	carFuel := jialun(1.2)
	//给car安排40升的油，但是人家本来剩下1.2单位加仑
	carFuel += tojialun(sheng(40))
	fmt.Printf("carFuel:%0.1f jialun",carFuel)

}