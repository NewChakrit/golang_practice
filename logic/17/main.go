package main

import "fmt"

func main() {
	fmt.Println(revertInput(1234))
}

func revertInput(n int) int {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
		fmt.Printf("reversed = %v\n", reversed)
		fmt.Printf("n = %v\n", n)
	}
	return reversed
}
