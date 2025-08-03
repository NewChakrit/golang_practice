package main

import "fmt"

/* Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.



Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.
Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].

Idea
1. เก็บค่าของ Roman ใน map[byte]int
2. วน string ทีละตัว
3. ถ้าค่าตัวนี้ < ตัวถัดไป → ลบ
   ถ้าไม่ → บวก
*/

func main() {
	fmt.Println(romanToInt("LVIII")) //58
	fmt.Println(romanToInt("IX"))    //9
	fmt.Println(romanToInt("XL"))    //40
}

func romanToInt(s string) int {
	total := 0
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
			total -= roman[s[i]]
		} else {
			total += roman[s[i]]
		}
	}

	//fmt.Println(total)
	return total
}

//s = "MCMXCIV"
//index: 0 1 2 3 4 5 6
//char:  M C M X C I V
//value: 1000 100 1000 10 100 1 5
//i = 0: M (1000) < C (100) ❌ → บวก 1000 → total = 1000
//i = 1: C (1000) < C (100) ❗ ลบ 100 → total = 900
//i = 3: M (1000) < X (10) ❌ → บวก 1000 → total = 1900
//i = 4: X (10) < C (100) ❗ ลบ 10 → total = 1890
//i = 5: C (100) < I (1)  ❌ → บวก 100 → total = 1990
//i = 6: I (1) < V (5) ❗ ลบ 1 → total = 1989
//i = 7: V (5) ❗ บวก 5 → total = 1994

//or
//func romanToInt(s string) int {
//	roman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
//	total := 0
//	for i := 0; i < len(s)-1; i++ {
//		if roman[s[i]] < roman[s[i+1]] {
//			total -= roman[s[i]]
//		} else {
//			total += roman[s[i]]
//		}
//	}
//	total += roman[s[len(s)-1]] // บวกตัวสุดท้ายทีหลัง
//	return total
//}
