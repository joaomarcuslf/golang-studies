package main

import (
	"fmt"
)

func findDuplicates(nums []int) int {
	slow := 0
	fast := 1

	for {
		if nums[slow] == nums[fast] {
			return nums[slow]
		}

		slow += 1
		fast += 2

		if fast >= len(nums) {
			fast = 0
		}

		if slow >= len(nums) {
			slow = 0
		}
	}
}

func main() {
	fmt.Println(findDuplicates([]int{2, 3, 4, 1, 5, 5}))
	fmt.Println(findDuplicates([]int{1, 2, 4, 4, 3, 3, 1, 5, 3}))
}
