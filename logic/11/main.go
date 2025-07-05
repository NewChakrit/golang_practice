package main

import "fmt"

/*วาดสามเหลี่ยมดาว (pyramid)
เขียนโปรแกรมรับ n แล้วแสดงสามเหลี่ยมดาวสูง n ชั้น เช่น n=4:
   *
  ***
 *****
*******
*/

func main() {
	m := 4

	for i := 1; i <= m; i++ {
		for j := m; j > i; j-- {
			fmt.Print(" ")
		}

		for p := 0; p < 2*i-1; p++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := 1; i <= m; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		for p := m; p > (2*i)-(m-1); p-- {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

//func main() {
//	pyramid(4)
//}
//
//func pyramid(n int) {
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= n-i; j++ {
//			fmt.Print(" ")
//		}
//
//		for k := 1; k <= 2*i-1; k++ {
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//}
