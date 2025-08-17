package main

import "fmt"

/*
You are given a sorted unique integer array nums.
A range [a,b] is the set of all integers from a to b (inclusive).
Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b


Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
Example 2:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

โจทย์
เรามี array ของจำนวนเต็ม ที่ เรียงจากน้อยไปมากอยู่แล้ว และ ไม่มีตัวเลขซ้ำ
เราต้องการแบ่งตัวเลขใน array นี้ออกเป็น ช่วง (range)
แต่ละช่วงต้อง ครอบคลุมตัวเลขที่ต่อเนื่องกัน เท่านั้น (ไม่มีหล่น ไม่มีแถม)
สุดท้าย เราต้องคืนรายการของช่วงทั้งหมด โดยเขียนในรูปแบบ string

การแสดงผลช่วง
ถ้าช่วงมีตัวเดียว เช่น [5,5] → เขียนเป็น "5"
ถ้าช่วงมีหลายตัว เช่น [2,4] (คือ 2,3,4) → เขียนเป็น "2->4"

1. ใช้ start เก็บค่า เริ่มต้นของช่วง
2. เช็คว่าเลขถัดไป ต่อเนื่องมั้ย (nums[i] != nums[i-1]+1)
3. บันทึกช่วง ทั้งกรณีเลขเดียว (start == nums[i-1]) หรือหลายตัว (start != nums[i-1])
4. หลัง loop ต้องบันทึกช่วงสุดท้ายเสมอ
*/

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))    // ["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // ["0","2->4","6","8->9"]
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := []string{}
	start := nums[0] // เริ่มช่วงแรก

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			// จบช่วง
			if start == nums[i-1] {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i] // เริ่มช่วงใหม่
		}
	}

	// บันทึกช่วงสุดท้าย
	if start == nums[len(nums)-1] {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, nums[len(nums)-1]))
	}

	return result
}
