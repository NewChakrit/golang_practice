package main

import "fmt"

/*
หาค่าซ้ำใน slice
เขียนฟังก์ชันรับ slice ของ int แล้ว return slice ของเลขที่ซ้ำกัน
*/

func main() {
	nums := []int{6, 7, 44, 66, 23, 8, 7, 9, 3, 4, 66, 4}

	fmt.Println(FindDuplicates(nums))
}

func FindDuplicates(nums []int) []int {
	counts := make(map[int]int) // ร้าง map ชื่อ counts มี key เป็นเลข และ value เป็นจำนวนครั้งที่พบ
	for _, num := range nums {
		counts[num]++ // ทุกครั้งมี่พบ num จะ + 1 value
	}

	var result []int
	for k, v := range counts {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}
