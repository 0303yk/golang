package main
//现在要解决一个问题:简化代码，如果两个类有相同方法的简化方法p331
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
//上面两个方法，下面
//  如果playlist方法体参数类型改为Tapecorde：
//  cannot use player (type TapePlayer) as type Tapecorde in argument to playlist
//虽然TapePlayer类和Tapecorde类都有Stop和play方法，但是实际操作不能改playlist方法中的参数
func playlist(device Tapecorde,songs []string){
	for _,song := range songs{
		device.Play(song)
	}
	device.Stop()
}

func main(){
	player := Tapecorde{	}
	mixtape := []string{"lily","look what you made me do","ear"}
	playlist(player,mixtape)
}