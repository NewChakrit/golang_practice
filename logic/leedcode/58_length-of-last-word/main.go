package main

import (
	"fmt"
	"strings"
)

/*
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.

*/

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

//func lengthOfLastWord(s string) int {
//	str := strings.TrimRight(s, " ")
//	count := 0
//
//	for i := len(str) - 1; i >= 0; i-- {
//		if str[i] != ' ' { // ต้องใช้ ' ' แทน " " เพราะsprS[i] เป็น byte (ตัวอักษรเดี่ยวในรูปแบบ uint8) " " (double quotes) เป็น string (ชุดของ byte) Go ไม่ยอมให้เปรียบเทียบ byte กับ string โดยตรง
//			count++
//		} else if count >= 1 {
//			break
//		}
//	}
//
//	return count
//}

func lengthOfLastWord(s string) int {
	sprS := strings.Split(s, "")
	count := 0
	for i := len(sprS) - 1; i >= 0; i-- {
		if sprS[i] != " " {
			count++
		} else if count >= 1 && sprS[i] == " " {
			break
		}
	}

	return count
}
