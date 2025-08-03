package main

/*
You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [2,3,0,1,4]
Output: 2

ให้ array ของจำนวนเต็ม nums ขนาด n ซึ่ง index เริ่มจาก 0 โดยคุณจะเริ่มต้นที่ตำแหน่ง nums[0]
แต่ละค่า nums[i] หมายถึง ระยะทางสูงสุดที่สามารถกระโดดไปข้างหน้าจากตำแหน่ง i
พูดอีกแบบคือ ถ้าคุณอยู่ที่ index i คุณจะสามารถกระโดดไปที่ nums[i + j] ได้ โดยที่:

0 <= j <= nums[i]

และ i + j < n

โจทย์ต้องการให้ คืนจำนวนครั้งที่กระโดดน้อยที่สุด เพื่อให้ไปถึง nums[n - 1] (ตำแหน่งสุดท้าย)
และโจทย์รับประกันว่าจาก input จะสามารถไปถึง nums[n - 1] ได้แน่นอน

ตัวอย่าง:
✅ Example 1:
Input: nums = [2,3,1,1,4]
Output: 2
คำอธิบาย: กระโดดจาก index 0 ไป index 1 จากนั้นกระโดดอีกทีจาก index 1 ไปถึง index สุดท้าย (index 4) รวมเป็น 2 ครั้ง

✅ Example 2:
Input: nums = [2,3,0,1,4]
Output: 2
*/
func main() {

}
