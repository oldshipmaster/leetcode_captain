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

func moveZeroes(nums []int)  {
	j:=0
	for _,v:= range nums {
		if v != 0 {
			nums[j] = v
			j++
		}
	}
	for i:=j; i< len(nums);i++ {
		nums[i] = 0
	}
}

func moveZeroes2(nums []int)  {
	for i,j:=0,0;i<len(nums);i++ {
		if nums[i] != 0 {
			nums[j],nums[i] = nums[i],nums[j]
			j++
		}
	}
}