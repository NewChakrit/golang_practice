package main

import "fmt"

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.



Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

มีอาร์เรย์จำนวนเต็มสองอัน nums1 และ nums2 ซึ่งถูกจัดเรียงในลำดับไม่ลดลง (non-decreasing order) แล้ว
และมีจำนวนเต็มสองค่า m และ n แทนจำนวนของสมาชิกจริง ๆ ใน nums1 และ nums2 ตามลำดับ

ให้ทำการรวม (merge) nums1 และ nums2 ให้เป็นอาร์เรย์เดียวที่ถูกจัดเรียงในลำดับไม่ลดลง

อาร์เรย์ที่จัดเรียงแล้วนี้ ไม่ต้อง return ออกมาจากฟังก์ชัน แต่ต้องจัดเก็บไว้ใน nums1 แทน โดย nums1 จะมีความยาวเท่ากับ m + n
ซึ่งสมาชิก m ตัวแรกจะเป็นข้อมูลจริงที่ต้องนำมารวม และสมาชิก n ตัวท้ายจะถูกกำหนดค่าเป็น 0 และให้มองข้ามได้ ส่วน nums2 จะมีความยาวเท่ากับ n
*/

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	fmt.Println(nums1)
}
