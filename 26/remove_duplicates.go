package _6

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Printf("%v", nums)
	return i + 1

}






















func removeDuplicates1(nums []int) int {
	i:=0
	for j := 1; j < len(nums); j++ {


		if nums[i] == nums[j] {
			continue
		} else {
			i++
			nums[i] = nums[j]
		}

	}
	//fmt.Printf("%v", nums[:index+1])
	return i + 1
}
