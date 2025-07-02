package main

import "fmt"

/*
ผลรวมเฉพาะเลขคู่
เขียนฟังก์ชันรับ n (int) แล้ว return ผลรวมของเลขคู่ตั้งแต่ 1 ถึง n

เช่น n=10 → 2+4+6+8+10=30
*/

func main() {
	fmt.Println(sumEven(10))
}

func sumEven(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			result += i
		}
	}

	return result
}
