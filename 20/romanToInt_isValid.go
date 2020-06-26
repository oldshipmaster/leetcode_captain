package _1

//https://leetcode-cn.com/problems/two-sum/

func isValid(s string) bool {
	v := []byte(s)
	j := len(v) - 1
	flag := true
	for i := 0; i < len(v); i++ {
		if v[i] +2 != v[j] {
			flag = false
			break
		}
		j--
	}
	if !flag {
		j = 0
		for k := 0; k < len(v); k++ {

			if v[j]+2!=v[j+1] {
				flag = false
				break
			}
			k = j
			k = k+2
			j = k
		}
	}
	return flag

}
