package _1

//https://leetcode-cn.com/problems/two-sum/

func reverseString(s []byte) {
	if len(s) < 1 {
		return
	}
	j := len(s) -1
	for i := 0; i < j; i++ {
		s[i],s[j] = s[j],s[i]
		j = j-1
	}

}

func reverseString1(s []byte) {
	if len(s) < 1 {
		return
	}
	j := len(s) -1
	for i := 0; i < len(s); i++ {
		if j < i {
			break
		}
		s[i],s[j] = s[j],s[i]
		j = j-1

	}

}

