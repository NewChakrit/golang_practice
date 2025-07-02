package main

import "fmt"

/* สร้าง slice ของเลข Fibonacci n ตัว
เขียนฟังก์ชันรับ n (int) แล้ว return slice ของตัวเลข Fibonacci n ตัวแรก

เช่น n=6 → [0 1 1 2 3 5]

Fibonacci sequence (ลำดับฟีโบนัชชี)
คือ ลำดับตัวเลขที่: ตัวเลขแต่ละตัว (ตั้งแต่ตัวที่ 3 เป็นต้นไป) เกิดจากผลบวกของ สองตัวก่อนหน้า

F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2)

0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
ตัวที่ 0 = 0
ตัวที่ 1 = 1
ตัวที่ 2 = 0+1 = 1
ตัวที่ 3 = 1+1 = 2
ตัวที่ 4 = 1+2 = 3
ตัวที่ 5 = 2+3 = 5
ตัวที่ 6 = 3+5 = 8
และต่อไปเรื่อย ๆ */

func main() {
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) []int {
	result := []int{}

	for i := 0; i < n; i++ {
		if i > 1 {
			result = append(result, result[i-2]+result[i-1])
		} else {
			result = append(result, i)
		}
	}

	return result
}
