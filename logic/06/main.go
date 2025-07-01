package main

import "fmt"

// หาค่า max ใน slice

func main() {
	fmt.Println(Max([]int{2, 4, 55, 7, 9, 15, 33, 77, 32, 145, 6, 2}))
}

func Max(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}

	result := nums[0] // handle ค่าใน nums ติดลบ
	for _, v := range nums {
		if v > result {
			result = v
		}
	}

	return result
}
