package main

import (
	"fmt"
	"sort"
)

/*เขียนฟังก์ชันรับ slice ของ string แล้วจัดกลุ่มคำที่เป็น anagram กันลงใน map
key เป็น string ที่เป็น sorted ตัวอักษรของคำ
value เป็น slice ของคำที่เป็น anagram กัน
เช่น input: []string{"eat", "tea", "tan", "ate", "nat", "bat"}
output:
map[string][]string{
  "aet": {"eat", "tea", "ate"},
  "ant": {"tan", "nat"},
  "abt": {"bat"},
}

rune คือ alias ของ int32
ใช้แทน character (ตัวอักษร) ใน Unicode

sort.Slice เป็นฟังก์ชันในแพ็กเกจ sort ของ Go
ใช้สำหรับ sort slice อะไรก็ได้ (ไม่จำกัดว่าต้องเป็น int หรือ string)
โดยให้เรากำหนดเงื่อนไขการเปรียบเทียบเอง
*/

func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	for k, v := range result {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

func groupAnagrams(words []string) map[string][]string {
	result := map[string][]string{}
	for _, word := range words {
		runes := []rune(word) // ทำให้กลายเป็น uncode
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j] // เรียง unicode
		})
		key := string(runes)

		result[key] = append(result[key], word)
		//fmt.Println(key)
	}

	return result
}
