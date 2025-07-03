package main

import (
	"fmt"
	"strings"
)

/*รับ string จากผู้ใช้ แล้วแสดงว่าแต่ละตัวอักษร (ไม่เว้น space) พบกี่ครั้ง
เช่น "banana" → b:1, a:3, n:2
(Hint: ใช้ loop + map)*/

func main() {
	var text string
	fmt.Print("Input some text: ")
	fmt.Scanln(&text)

	result := countEachString(text)
	for k, v := range result {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func countEachString(s string) map[string]int {
	str := map[string]int{}
	strList := strings.Split(s, "")
	for _, v := range strList {
		str[v]++
	}

	return str
}
