package main

import "fmt"

func main() {
	// 1. ลบ element ที่ index i
	i := 3
	nums1 := []int{2, 4, 6, 8, 10, 12, 14, 16}
	nums1 = append(nums1[:i], nums1[i+1:]...)
	fmt.Println(nums1) // [2 4 6 10 12 14 16]

	// 2. Reverse slice
	nums2 := []int{2, 4, 6, 8, 10, 12, 14, 16}
	for i, j := 0, len(nums2)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i:%v , j:%v\n", i, j)
		fmt.Printf("num[i]:%v , num[j]:%v\n", nums2[i], nums2[j])
		nums2[i], nums2[j] = nums2[j], nums2[i]
	}
	fmt.Println(nums2)

	// 3. Unique (ลบค่าซ้ำ)
	nums3 := []int{2, 10, 8, 8, 10, 10, 10, 12, 14, 14}
	m := map[int]bool{}

	var result []int
	for _, v := range nums3 {
		if !m[v] { // ถ้า map m ยัง ไม่มีค่า v (หรือยังเป็น false/ค่าเริ่มต้น)
			m[v] = true
			result = append(result, v)
		}
	}
	//
	fmt.Println(result)

	// find min max
	nums4 := []int{2, 4, 6, 8, 10, 12, 14, 16}
	var min, max = nums4[0], nums4[0]
	for _, v := range nums4 {
		if min < v {
			min = v
		} else {
			max = v
		}

	}
	fmt.Printf("min: %v, max: %v\n", min, max)
}
