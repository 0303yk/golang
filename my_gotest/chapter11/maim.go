package main
//有相同方法的两种不同类型P322
import "fmt"
type TapePlayer struct{
	Batteries string
}
func (t TapePlayer) Play(song string){
	fmt.Println("playing",song)
}
func (t TapePlayer) Stop(){
	fmt.Println("stop")
}
type Tapecorde struct{
	Microphone int
}
func (t Tapecorde) Play(song string){
	fmt.Println("playing",song)
}
func (t Tapecorde) Record(){
	fmt.Println("Recording")
}
func (t Tapecorde) Stop(){
	fmt.Println("stop")
}
func playlist(device Tapecorde,songs []string){
	for _,song := range songs{
		device.Play(song)
	}
	device.Stop()
}

func main(){
	player := TapePlayer{	}
	mixtape := []string{"lily","look what you made me do","ear"}
	playlist(player,mixtape)
}