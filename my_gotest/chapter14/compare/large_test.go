package compare

import (
	"fmt"
	"testing"
)
//要优化test函数
func TestFirstLarger(t *testing.T){
	want := 2
	got := Lager(2,1)
	if got != want {
		t.Error(errorString(2,1,got,want))
	}
}


func TestSecongLarger(t *testing.T){
	want := 8
	got := Lager(4,8)
	if got != want {
		t.Error(errprString(4,8,got,want))
	}
}
//下面是helper函数
func errprString(a int , b int , want int)string {
	return fmt.Sprintf("Larger(%d,%d) = %d ,want %d",a,b,got,want)
}
