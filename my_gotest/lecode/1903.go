package lecode
//https://leetcode.cn/problems/largest-odd-number-in-string/
func largestOddNumber(num string) string {

	l := len(num)

	for i := l - 1; i >= 0; i-- {
		if ((num[i] - '0') & 1) == 1 {
			return num[0:i+1]
		}
	}

	return ""

}
