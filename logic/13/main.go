package main

import "fmt"

/*พิมพ์ตารางสูตรคูณ 1-12
เขียนโปรแกรม loop แสดงตารางสูตรคูณตั้งแต่ 1 ถึง 12*/

func main() {
	for i := 1; i <= 12; i++ {
		fmt.Println("multiplication table ", i)
		for j := 1; j <= 12; j++ {
			mul := i * j
			fmt.Printf("%v x %v = %v\n", i, j, mul)
		}
	}
}
