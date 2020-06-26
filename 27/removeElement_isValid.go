package _1

//https://leetcode-cn.com/problems/two-sum/

//拷贝覆盖
func removeElement(nums []int, val int) int {
	c:=0
	for _,v := range nums{
		if v != val {
			nums[c] = v
			c++
		}
	}
	return c
}