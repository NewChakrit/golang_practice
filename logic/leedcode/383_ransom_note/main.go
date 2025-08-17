package main

import "fmt"

/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.

Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true

กำหนดให้มีสตริงสองตัวคือ ransomNote และ magazine ให้คืนค่า true หากสามารถสร้าง ransomNote โดยใช้ตัวอักษรจาก magazine และคืนค่า false หากไม่สามารถทำได้
ตัวอักษรแต่ละตัวใน magazine สามารถใช้ได้เพียงครั้งเดียวใน ransomNote

rune

คือ ตัวอักษร 1 ตัวใน Unicode (เป็น alias ของ int32)
เวลาเราเขียน for _, ch := range s { ... } ตัวแปร ch จะเป็น rune
แบบนี้ Go จะตัด string ให้เป็น ตัวอักษรจริง ๆ (ไม่ใช่ byte)
เหมาะสำหรับรองรับภาษาโลก เช่น "สวัสดี"
s := "สวัสดี"

// วนด้วย byte
for i := 0; i < len(s); i++ {
    fmt.Printf("%d ", s[i])
}
// output: ตัวเลขยาว ๆ หลายค่า (เพราะ 1 ตัวอักษรไทยกินหลาย byte)

// วนด้วย rune
for _, ch := range s {
    fmt.Printf("%c ", ch)
}
// output: ส วั ส ดี   (ตัดเป็นตัวอักษรถูกต้อง)

*/

func main() {
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "ab"))  // false
	fmt.Println(canConstruct("aa", "aab")) // true
}

func canConstruct(ransomNote string, magazine string) bool {
	counts := make(map[rune]int)

	for _, ch := range magazine {
		counts[ch]++
	}

	for _, ch := range ransomNote {
		if counts[ch] == 0 {
			return false
		}
		counts[ch]--
	}

	return true
}

/* ใช้ byte แทน rune
func canConstruct(ransomNote string, magazine string) bool {
	// สร้าง array ขนาด 26 สำหรับ a-z
	counts := [26]int{}

	for i := 0; i < len(magazine); i++ {
		ch := magazine[i] - 'a' // หาตำแหน่งใน array เช่น 'c' - 'a' = 2
		counts[ch]++
	}

	// ตรวจสอบตัวอักษรจาก ransomNote
	for i := 0; i < len(ransomNote); i++ {
		ch := ransomNote[i] - 'a'
		if counts[ch] == 0 {
			return false
		}
		counts[ch]--
	}

	return true
}

ทำไมต้อง - 'a' ?

เพราะเราต้องการแปลงตัวอักษร a–z ให้กลายเป็นเลข index ของ array [26]int
'a' - 'a' = 97 - 97 = 0 → เก็บในช่องที่ 0
'b' - 'a' = 98 - 97 = 1 → เก็บในช่องที่ 1
'c' - 'a' = 99 - 97 = 2 → เก็บในช่องที่ 2

ทำไมใช้ array ได้เลย?

เพราะเรารู้แน่นอนว่า มีแค่ 26 ตัวอักษร
เลยสามารถใช้ array [26]int แทน map เพื่อเก็บจำนวนได้

counts[0] = จำนวน 'a'
counts[1] = จำนวน 'b'
…
counts[25] = จำนวน 'z'
*/
