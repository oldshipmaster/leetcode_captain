package _6

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	nums := []int{1,2,3}
	sum := removeDuplicates1(nums)
	fmt.Printf("%d", sum)
}
