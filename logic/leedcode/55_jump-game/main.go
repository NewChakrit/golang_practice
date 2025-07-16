package main

import "fmt"

/*
You are given an integer array nums. You are initially positioned
at the array's first index, and each element in the array represents your maximum jump length at that position.
Return true if you can reach the last index, or false otherwise.

Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

คุณได้รับอาร์เรย์ของจำนวนเต็ม nums
คุณจะเริ่มต้นอยู่ที่ตำแหน่งแรกของอาร์เรย์
โดยที่แต่ละค่าใน nums แทนความยาวสูงสุดที่คุณสามารถกระโดดได้จากตำแหน่งนั้น

ให้เขียนฟังก์ชันเพื่อตรวจสอบว่า
คุณสามารถกระโดดไปถึงตำแหน่งสุดท้ายของอาร์เรย์ได้หรือไม่
ถ้าได้ ให้คืนค่า true
ถ้าไม่ได้ ให้คืนค่า false

✅ Example 1:
Input: nums = [2,3,1,1,4]
Output: true
อธิบาย:

กระโดดจาก index 0 ไป index 1 (กระโดด 1 ช่อง)

จาก index 1 กระโดดต่อไป index 4 (กระโดด 3 ช่อง)

ไปถึงตำแหน่งสุดท้ายได้

✅ Example 2:
Input: nums = [3,2,1,0,4]
Output: false
อธิบาย:

ไม่ว่าคุณจะเลือกกระโดดยังไง สุดท้ายจะต้องมาติดอยู่ที่ index 3

index 3 มีค่ากระโดดได้สูงสุดแค่ 0 → ทำให้ไปต่อไม่ได้

concept:
หาว่าเรากระโดดได้ไกลสุดได้เท่าไหร่แล้งเก็บค่าไว้ ลูปไปเรื่อยๆ แล้วลองเทียบค่าดู
*/

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {

	maxReach := 0
	for i := 0; i < len(nums)-1; i++ {
		if i > maxReach {
			return false
		} else {
			maxReach = max(maxReach, i+nums[i])
		}
	}
	return true
}
