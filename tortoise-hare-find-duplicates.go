package main

import (
	"fmt"
)

func findDuplicates(nums []int) int {
	slow := 0
	fast := 1
	aux1 := 0
	aux2 := 0

	for {
		if nums[slow] == nums[fast] && slow != fast {
			return nums[slow]
		}

		slow += 1
		fast += 2

		if fast >= len(nums) {
			aux1 += 1
			
			if (aux1) >= len(nums) {
				aux1 = 0
			}
			
			fast = aux1
		}

		if slow >= len(nums) {
			aux2 += 1
			
			if (aux2) >= len(nums) {
				aux2 = 0
			}
			
			slow = aux2
		}
	}
}

func main() {
	fmt.Println(findDuplicates([]int{2, 3, 4, 1, 5, 5}))
	fmt.Println(findDuplicates([]int{2, 4, 3, 3, 1, 5, 3}))
	fmt.Println(findDuplicates([]int{2, 3, 5, 6, 1, 1}))
}
