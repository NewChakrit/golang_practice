package main

import "fmt"

/*
Factorial (recursive & loop)
เขียนฟังก์ชันหาค่า factorial ด้วย recursive และ loop

factorial ของจำนวนเต็มบวก n (เขียนว่า n!) คือ ผลคูณของตัวเลขตั้งแต่ 1 ถึง n
เช่น 5! = 5 × 4 × 3 × 2 × 1 = 120
    3! = 3 × 2 × 1 = 6
    0! = 1 (นิยามพิเศษ)
recursive คือการเขียนฟังก์ชันที่ เรียกตัวเอง
เช่น factorial: n! = n × (n-1)!
*/

func main() {

	fmt.Println(factorial(5))
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
