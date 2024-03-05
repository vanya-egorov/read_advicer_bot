package main

import "fmt"

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			if nums[j] == nums[j+1] {
				return true
			}
		}
	}
	return false
}

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println(cap(mySlice))
}
