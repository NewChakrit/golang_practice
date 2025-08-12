package main

import (
	"fmt"
	"strings"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

หา prefix ที่ยาวที่สุด (ตัวอักษรที่อยู่ต้นเหมือนกันในทุก string)

ใช้คำแรกเป็น ฐาน (base)

ไล่เช็คทีละตัวอักษรกับคำอื่นๆ

ถ้าตัวอักษรใดไม่ตรงกัน ให้ break

ส่วนที่เหลือคือ common prefix
*/

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, v := range strs {
		for !strings.HasPrefix(v, prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
