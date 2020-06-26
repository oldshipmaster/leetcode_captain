package _1

import (
	"strings"
)

//https://leetcode-cn.com/problems/two-sum/

func reverse(s []byte) []byte {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j = j - 1
	}
	return s
}

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	var ss []string
	for _, word := range split {
		s := []byte(word)
		str := string(reverse(s))
		ss = append(ss, str)
	}
	return strings.Join(ss, " ")
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}
	max_key :=0
	max_val := 0
	for k, v := range m {
		if v > max_val {
			max_val = v
			max_key = k
		}
	}
	return max_key

}
