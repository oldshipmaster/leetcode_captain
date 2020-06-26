package _1

//https://leetcode-cn.com/problems/two-sum/

func removeElement(nums []int, val int) int {
	j:=0
	i:=0
	for i < len(nums){
		if nums[j] == val {
			i = i+1
			nums[j]= nums[i]
			i = j
		} else {
			j = j+1
		}
	}
	return len(nums)
}
