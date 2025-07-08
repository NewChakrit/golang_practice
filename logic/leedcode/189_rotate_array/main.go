package main

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

ให้ อาร์เรย์ของจำนวนเต็ม nums
ให้คุณ หมุนอาร์เรย์ไปทางขวา (right) จำนวน k ครั้ง
โดยที่ k เป็นจำนวนเต็ม ≥ 0
*/

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // กัน k > n

	tmp := make([]int, k)
	copy(tmp, nums[n-k:]) // เก็บ k ตัวท้าย

	// เลื่อน n-k ตัวแรกไปขวา k ตำแหน่ง
	for i := n - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	// ใส่ tmp กลับไปที่ด้านหน้า
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}
