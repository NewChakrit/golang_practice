package main

import "fmt"

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

คุณได้รับอาร์เรย์ prices ซึ่ง prices[i] คือราคาหุ้นในวันที่ i
คุณต้องการหากำไรสูงสุดจากการซื้อหุ้นเพียง 1 ครั้ง แล้วขายในวันถัดไป (ในอนาคต)

ให้คุณหากำไรสูงสุดที่สามารถทำได้จากการทำธุรกรรมนี้
ถ้าทำกำไรไม่ได้เลย ให้คืนค่าเป็น 0

📌 ตัวอย่าง:
Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
อธิบาย: ซื้อในวันที่ 2 (ราคา 1) แล้วขายในวันที่ 5 (ราคา 6) → ได้กำไร 6−1=5

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
อธิบาย: ในกรณีนี้ไม่มีวันที่จะซื้อแล้วขายได้กำไร → กำไรสูงสุด = 0
*/

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	// prices = [7,1,5,3,6,4]
	// หาค่าที่น้อยที่สุดก่อน คือ วันที่ซื้อ
	// หาตำแหน่งของค่าน้อยสุด คือ วันที่ซื้อ - 1
	// หาค่าที่มากที่สุด เริ่มนับจากตำแหน่งของค่าน้อยสุดที่หาได้ แต่ถ้าหาไม่ได้ ให้ set เป็น 0

	buy := prices[0]
	mProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		}

		if mProfit < prices[i]-buy {
			mProfit = prices[i] - buy
		}
	}

	return mProfit
}
