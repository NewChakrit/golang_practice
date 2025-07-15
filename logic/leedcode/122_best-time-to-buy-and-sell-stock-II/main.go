package main

import "fmt"

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
Find and return the maximum profit you can achieve.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
Example 2:

Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.
Example 3:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.

คุณได้รับอาร์เรย์จำนวนเต็ม prices ซึ่ง prices[i] คือราคาหุ้นในวันที่ i
ในแต่ละวัน คุณสามารถตัดสินใจที่จะซื้อ และ/หรือ ขายหุ้นได้
คุณสามารถถือหุ้นได้อย่างมากที่สุดแค่ 1 หุ้นในเวลาเดียวกัน
แต่คุณสามารถซื้อแล้วขายทันทีในวันเดียวกันได้

ให้หากำไรสูงสุดที่คุณจะสามารถทำได้จากการซื้อ–ขายหุ้นตามเงื่อนไขข้างต้น

📌 ตัวอย่าง:
✅ Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 7
อธิบาย:

ซื้อวัน 2 (ราคา 1) → ขายวัน 3 (ราคา 5) → กำไร 4

ซื้อวัน 4 (ราคา 3) → ขายวัน 5 (ราคา 6) → กำไร 3
รวมกำไร = 4+3=7

✅ Example 2:
Input: prices = [1,2,3,4,5]
Output: 4
อธิบาย:

ซื้อวัน 1 (ราคา 1) → ขายวัน 5 (ราคา 5) → กำไร 4

✅ Example 3:
Input: prices = [7,6,4,3,1]
Output: 0
อธิบาย:

ไม่มีโอกาสทำกำไร → กำไรสูงสุด = 0
*/

// 1. ดูว่าวันข้างหน้าถัดไป 1 วันมีค่ามากกว่ามั้ย
// 2. ถ้ามากกว่าให้ขายในวันถัดไป แล้ว + กำไรไว้สะสมไว้

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 7
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))    // 4
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}

func maxProfit(prices []int) int {
	mProfit := 0
	for i := 0; i < len(prices); i++ {
		if i < len(prices)-1 && prices[i+1] > prices[i] {
			mProfit += prices[i+1] - prices[i]
		}
	}

	fmt.Printf("mProfit: %v", mProfit)

	return mProfit
}
