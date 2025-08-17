package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.

มี สอง string:
haystack → ข้อความหลัก (เหมือนกองฟาง)
needle → ข้อความที่เราต้องการค้นหา (เหมือนเข็มในกองฟาง)

ให้เราหา ตำแหน่งเริ่มต้น (index) ของ needle ใน haystack
ถ้าไม่เจอ → return -1
*/

func main() {
	strStr("sadbutsad", "sad")
	strStr("leetcode", "leeto")
}

//
//func strStr(haystack string, needle string) int {
//	result := strings.Index(haystack, needle)
//
//	return result
//}

//func strStr(haystack string, needle string) int {
//	result := strings.Index(haystack, needle)
//
//	return result
//}

func strStr(haystack string, needle string) int {
	haySpl := strings.Split(haystack, "")
	needSpl := strings.Split(needle, "")

	for i := 0; i <= len(haySpl)-len(needSpl); i++ {
		j := 0
		for ; j < len(needSpl); j++ {
			if haySpl[i+j] != needSpl[j] {
				break
			}
		}
		if j == len(needSpl) {
			return i
		}
	}

	return -1
}

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	cleanStr := reg.ReplaceAllString(s, "")
	cleanStr = strings.ToLower(cleanStr)

	fmt.Println("cleanStr =", cleanStr)
	fmt.Println(len(cleanStr))

	//for i := 0, j := len(cleanStr; i < len(cleanStr)/2; i = i+1, j = j-1 {
	//	if cleanStr[i] != cleanStr[j]{
	//		return false
	//	}
	//}

	return true
}
