package _1

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1,2,3,4,6}
	target := 11
	m := twoSum4(nums, target)
	fmt.Printf("%v",m)

}