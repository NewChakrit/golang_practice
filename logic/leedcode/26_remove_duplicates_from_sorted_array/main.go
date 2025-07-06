package main

import "fmt"

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
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

ให้ nums เป็นอาร์เรย์จำนวนเต็มที่ถูกจัดเรียงแล้วในลำดับไม่ลดลง (non-decreasing order)
ให้ลบค่าที่ซ้ำกันออก ในที่เดิม (in-place) เพื่อให้แต่ละค่าที่เหลือเป็นค่าที่ไม่ซ้ำกัน และแต่ละค่าปรากฏเพียงครั้งเดียว
ลำดับของสมาชิกที่เหลือ ต้องเหมือนเดิมตามที่เคยปรากฏใน nums ตอนแรก
จากนั้น ให้ return จำนวนค่าที่ไม่ซ้ำกันใน nums

สมมติว่าจำนวนค่าที่ไม่ซ้ำกันคือ k
เพื่อให้ผ่านการตรวจ คุณต้องทำตามนี้:

แก้ไข nums ให้ k ตำแหน่งแรกเป็นค่าที่ไม่ซ้ำกัน (และเรียงลำดับตามที่เคยอยู่ใน nums ตอนแรก)

ส่วนที่เหลือของ nums หลังจาก k ตำแหน่งแรก ไม่สำคัญ (จะเป็นค่าอะไรก็ได้)

return ค่า k
*/

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	p1 := 0
	for i, v := range nums {
		fmt.Printf("v: %v\n", v)
		// fmt.Printf("p1: %v\n",p1)
		fmt.Printf("nums[p1]: %v, p1: %v\n", nums[p1], p1)

		if v != nums[p1] && i > 0 {
			nums[p1+1] = v
			p1++
		}
		fmt.Println(nums)
	}

	return p1 + 1
}

func removeDuplicatesS2(nums []int) int {
	p1 := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[p1] {
			p1++
			nums[p1] = nums[i]
		}
	}

	return p1 + 1
}

/* Logic
เช็ค p1 ถ้าไม่เท่ากับ num[i] (ซึ่ง loop ค่าไปด้านขวาเรื่อยๆ)
ให้ชี้ค้าไปที่ค่า unique ถัดไปแล้วแทนค่า แล้วก็ทำซ้ำ
*/
