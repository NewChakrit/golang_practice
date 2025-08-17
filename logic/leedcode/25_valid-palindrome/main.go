package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

หากหลังจากแปลงอักษรตัวพิมพ์ใหญ่ทั้งหมดเป็นตัวพิมพ์เล็กและลบอักขระที่ไม่ใช่ตัวอักษรและตัวเลขทั้งหมดแล้ว วลีดังกล่าวอ่านได้เหมือนกันทั้งด้านหน้าและด้านหลัง อักขระตัวอักษรและตัวเลขประกอบด้วยตัวอักษรและตัวเลข
*/

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	cleanStr := reg.ReplaceAllString(s, "")
	cleanStr = strings.ToLower(cleanStr)

	for i, j := 0, len(cleanStr)-1; i < len(cleanStr)/2; i, j = i+1, j-1 {
		if cleanStr[i] != cleanStr[j] {
			return false
		}
	}

	return true
}
