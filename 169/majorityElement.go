package _1


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
