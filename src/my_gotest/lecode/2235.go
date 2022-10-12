package lecode
import "fmt"
func sum(num1 int, num2 int) int {
    for num2 != 0{
        t := num1
        num1 ^= num2
        num2 = (t & num2) << 1
    }
    return num1
}

func main(){
    fmt.Println(sum(4,5))
}
//作者：hzh512
//链接：https://leetcode.cn/problems/add-two-integers/solution/wei-yun-suan-by-hzh512-m7so/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

