package main

import (
	"fmt"
	"strconv"
)

/*  FizzBuzz
เขียนฟังก์ชันที่รับค่า n แล้ว return slice ของ string
ตามกติกา:

ถ้าหาร 3 ลงตัว → ใส่ "Fizz"

ถ้าหาร 5 ลงตัว → ใส่ "Buzz"

ถ้าหารทั้ง 3 และ 5 ลงตัว → ใส่ "FizzBuzz"

ถ้าไม่ใช่ → ใส่เลขเดิมเป็น string */

func main() {
	var num int

	for {
		fmt.Print("Input number: ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		fmt.Println(FizzBuzz(num))
	}
}

func FizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}
