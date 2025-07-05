package main

import "fmt"

/*
Longest Substring Without Repeating Characters
เขียนฟังก์ชันรับ string แล้ว return ความยาวของ substring ที่ไม่มีตัวอักษรซ้ำกัน

ใช้ map ช่วย track ตัวอักษรและตำแหน่ง

เช่น input: "abcabcbb" → output: 3 (substring: "abc")
*/

func main() {
	substringWithoutRepeatingCharacters("pwwkew")
}

func substringWithoutRepeatingCharacters(s string) {
	charIndex := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, c := range s {
		if lastIndex, found := charIndex[c]; found && lastIndex >= start {
			// พบตัวซ้ำ และอยู่ใน window ตอนนี้ → เลื่อน start
			start = lastIndex + 1
		}
		// บันทึกตำแหน่งล่าสุดของตัวอักษรนี้
		charIndex[c] = i
		// คำนวณความยาว substring ตอนนี้
		length := i - start + 1
		if length > maxLength {
			maxLength = length
		}
	}
	fmt.Println(maxLength)
}

//func substringWithoutRepeatingCharacters(s string) {
//	strs := strings.Split(s, "")
//	m := map[string]bool{}
//	count := 0
//	for _, str := range strs {
//		if !m[str] {
//			m[str] = true
//			count++
//		}
//	}
//	fmt.Println(count)
//}
