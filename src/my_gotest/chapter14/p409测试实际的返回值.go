//测试实际的返回值，要有main包，同时有Test文件嗷
//P409页有更好的解答：在同一个包中，其中的文件名
// 为**_test.go，书写Test为开头的函数 用if判定结果
// （结果若与预期不符，可打印该错误） 这也是java许多包里
// 都有测试类的原因吧,若想更精确，可以使用Errorf（下面是Error，
// 只记录会出错）坑：但是为啥同指针呢？，
package chapter14

import (
	"fmt"
	"strings"
	"testing"
	)
func Joinwithcommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1],",")
	result += " and "
	result += phrases[len(phrases)-1]
	return result
}

func TestTwoElements(t *testing.T){
	//list是自己书写的示例，实现的结果也可以预见
	list := []string{"apple","orange"}
	if Joinwithcommas(list) != "apple and orange"{
	t.Error("no test written yet")
	}
}

func TestThreeElements(t *testing.T){
	//list是自己书写的示例，实现的结果也可以预见
	list := []string{"apple","orange","pear"}
	if Joinwithcommas(list) != "apple,orange and pear"{
		t.Error("no test written yet")
	}
}

func main() {
	//fmt.Println(strings.Join([]string{"state","of","the","art"},"-"))
	phrases := []string{"mum","dad"}
	fmt.Println("a photo of ",Joinwithcommas(phrases))
}
