package main

import "fmt"

/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                            // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

Example 1:

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).

ให้ nums เป็นอาร์เรย์ของจำนวนเต็ม และให้ val เป็นจำนวนเต็มอีกค่า
ให้ลบทุกค่าที่เท่ากับ val ออกจาก nums ในที่เดิม (in-place)
ลำดับของสมาชิกที่เหลืออาจเปลี่ยนแปลงได้
จากนั้นให้ return จำนวนสมาชิกใน nums ที่ไม่เท่ากับ val

สมมติว่าจำนวนสมาชิกที่ไม่เท่ากับ val คือ k
เพื่อให้ผ่านการตรวจ คุณต้องทำตามนี้:

แก้ไข nums ให้ k ตำแหน่งแรก เป็นค่าที่ไม่เท่ากับ val

ส่วนที่เหลือของ nums หลังจาก k ตำแหน่งแรก ไม่สำคัญ (จะเป็นอะไรก็ได้)

return ค่า k

Custom Judge (ตัวตรวจ):
ตัวตรวจจะทดสอบโค้ดของคุณด้วยโค้ดแบบนี้:

int[] nums = [...]; // อาร์เรย์ input
int val = ...; // ค่าที่ต้องลบ
int[] expectedNums = [...]; // คำตอบที่ถูกต้อง, จะไม่มีค่าเท่ากับ val และถูก sort ไว้

int k = removeElement(nums, val); // เรียกฟังก์ชันของคุณ

assert k == expectedNums.length;
sort(nums, 0, k); // sort k ตำแหน่งแรกของ nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}
ถ้า assert ทั้งหมดผ่าน แสดงว่าคำตอบถูกต้อง

ตัวอย่าง 1:
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,,]
อธิบาย: ต้อง return k = 2 และสองตำแหน่งแรกเป็น 2
หลังจากตำแหน่ง k จะเป็นค่าอะไรก็ได้ (แสดงเป็น _)

ตัวอย่าง 2:
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,,,_]
อธิบาย: ต้อง return k = 5 และห้าตำแหน่งแรกเป็นค่าที่ไม่เท่ากับ 2 (ลำดับไม่สำคัญ)
หลังจากตำแหน่ง k จะเป็นค่าอะไรก็ได้ (แสดงเป็น _)
*/

func main() {
	fmt.Println(removeElement2([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement2(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}
	fmt.Println(nums)
	return n
}

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v != val {
			// fmt.Printf("v: %v\n", v)
			nums[index] = v
			// fmt.Println("-------------")
			index++
		}
	}

	// nums = nums[:index]
	fmt.Println(nums)

	return index
}
