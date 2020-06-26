package _1

import "strconv"

//https://leetcode-cn.com/problems/two-sum/

func isPalindrome(x int) bool {
	s := []byte(strconv.Itoa(x))
	j := len(s) -1
	for i := 0; i < len(s); i++ {
		if s[i] == s[j] {
			j--
		} else {
			return false
		}
	}
	return true
}
