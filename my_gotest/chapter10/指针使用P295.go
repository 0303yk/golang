package main4
//指针使用P295
import "fmt"

type Coordinates struct{
	Latitude float64
	Longitude float64
}

func (c *Coordinates) setlatitude (latitude float64){
	c.Latitude = latitude
}
func (c *Coordinates) setlongitude (longitude float64){
	c.Longitude =longitude
}


func main(){
	co :=Coordinates{}
	co.setlatitude(37.42)
	co.setlongitude(-122.08)
	fmt.Println(co)
}