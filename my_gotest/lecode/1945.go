package lecode
//https://leetcode.cn/problems/sum-of-digits-of-string-after-convert/

func getLucky(s string, k int) (sum int) {
	for _, b := range s {
		b -= 'a' - 1
		sum += int(b/10 + b%10)
	}
	for k--; k > 0 && sum > 9; k-- { // sum < 10 时可以提前退出
		s := sum
		for sum = 0; s > 0; s /= 10 {
			sum += s % 10
		}
	}
	fmt.Println(sum)
	return
}
func main() {
	getLucky("zk",2)
}
