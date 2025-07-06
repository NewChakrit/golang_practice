package main

import "fmt"

/*
Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:
The judge will test your solution with the following code:
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

Example 1:

Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

ให้ nums เป็นอาร์เรย์จำนวนเต็ม ที่ถูกจัดเรียงในลำดับไม่ลดลง (non-decreasing order)
ให้ลบค่าที่ซ้ำกันบางส่วนออก ในที่เดิม (in-place)
โดยให้แต่ละค่าที่เหลือ ปรากฏได้ไม่เกินสองครั้ง
ลำดับของสมาชิกต้องเหมือนเดิมตามที่เคยปรากฏใน nums

เนื่องจากในบางภาษา เราไม่สามารถเปลี่ยนความยาวของอาร์เรย์ได้
ดังนั้นให้เราทำให้ผลลัพธ์สุดท้าย อยู่ใน nums ส่วนแรก (k ตำแหน่งแรก)
โดยที่ k คือตัวนับจำนวนสมาชิกหลังจากลบ duplicates ตามเงื่อนไขแล้ว
ค่าที่เกินจาก k ไปไม่ต้องสนใจ (จะเป็นค่าอะไรก็ได้)

ต้องห้ามใช้พื้นที่อาร์เรย์ใหม่เพิ่ม (extra space)
ต้องแก้ไข nums ในที่เดิม (in-place) โดยใช้หน่วยความจำเพิ่มเติม O(1)

Custom Judge (ตัวตรวจ):
ตัวตรวจจะทดสอบโค้ดของคุณโดยใช้โค้ดประมาณนี้:

int[] nums = [...]; // อาร์เรย์ input
int[] expectedNums = [...]; // ผลลัพธ์ที่ถูกต้อง (ความยาว k)

int k = removeDuplicates(nums); // เรียกฟังก์ชันของคุณ

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
ถ้า assert ทั้งหมดผ่าน แสดงว่าคำตอบถูกต้อง

✏️ ตัวอย่าง:
Example 1:
Input: nums = [1,1,1,2,2,3]
Output: k=5 → nums = [1,1,2,2,3,_]
อธิบาย: ค่าแต่ละค่าจะปรากฏได้ไม่เกินสองครั้ง

Example 2:
Input: nums = [0,0,1,1,1,1,2,3,3]
Output: k=7 → nums = [0,0,1,1,2,3,3,,]
อธิบาย: ค่า 1 เดิมมี 4 ตัว จะเหลือแค่ 2 ตัว

*/

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

func removeDuplicates(nums []int) int {
	// [1,1,1,2,2,3] -> input
	// [1,1,2,2,3,_] -> tobe

	// if input ยาว < 2 → ไม่มี index[0] ให้ใช้
	if len(nums) <= 2 {
		return len(nums)
	}

	n := 2
	for i := 2; i < len(nums); i++ {
		if nums[n-2] != nums[i] {
			nums[n] = nums[i]
			n++
		}
		fmt.Printf("i:%v , n:%v, nums: %v\n", i, n, nums)
	}

	fmt.Printf("nums: %v\n", nums)
	fmt.Println(n)

	return n
}
