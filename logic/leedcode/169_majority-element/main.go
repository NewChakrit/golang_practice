package main

/*
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2

ให้ อาร์เรย์ nums ขนาด n
ให้คุณ คืนค่าตัว majority element

Majority element คือค่าที่ปรากฏในอาร์เรย์ บ่อยกว่าจำนวน ⌊n / 2⌋ ครั้ง
(และคุณสามารถสมมติได้ว่า อาร์เรย์นี้จะมี majority element อยู่แน่นอน)
ตัวอย่างที่ 1:

Input: nums = [3,2,3]
Output: 3
เพราะเลข 3 ปรากฏ 2 ครั้ง ซึ่งมากกว่า ⌊3/2⌋ = 1 ครั้ง

ตัวอย่างที่ 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2
เพราะเลข 2 ปรากฏ 4 ครั้ง ซึ่งมากกว่า ⌊7/2⌋ = 3 ครั้ง
*/

func main() {

}

func majorityElement(nums []int) int {
	// โจทย์นี้ ยังไงมันก็มี candidate อยู่แล้ว
	// มีตัวแปร candidate เก็บค่า “ตัวเต็ง”
	// มีตัวแปร count เก็บ “น้ำหนักเสียง” ของตัวเต็ง
	// loop อ่านค่าทีละตัว
	// ถ้า count == 0 → ตั้ง candidate = ค่านี้ ()
	// ถ้า nums[i] == candidate → count++
	// ถ้า nums[i] != candidate → count--
	// สุดท้าย candidate จะเป็น majority element

	p := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 { // อันนี้สำคัญ คือ candidate ไหนไม่มากจริงๆ มันจะไปนับ candidate เพราะโดน count-- จนเหลือ 0
			p = nums[i]
		}

		if nums[i] == p {
			count++
		} else {
			count--
		}
	}

	return p
}

// Boyer–Moore majority vote algorithm
