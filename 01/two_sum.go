package _1

//https://leetcode-cn.com/problems/two-sum/

func TwoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if i2, ok := m[target-v]; ok {
			return []int{i2, i}
		}
		m[v] = i
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[i] = v
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-v {
				return []int{j, i}
			}
		}
	}
	return nil
}

func twoSum4(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left, right}
		}
		if nums[left]+nums[right] < target {
			left++
		}
		if nums[left]+nums[right] > target {
			right--
		}
	}
	return []int{-1, -1}

}