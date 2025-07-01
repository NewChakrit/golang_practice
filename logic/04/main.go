package main

import (
	"fmt"
	"strings"
)

func main() {

	var inputStr string
	for {
		fmt.Print("Input some text: ")
		fmt.Scan(&inputStr)
		if strings.ToLower(inputStr) == "exit" {
			break
		}

		fmt.Println(IsPalindrome(inputStr))
	}

}

func IsPalindrome(s string) bool {
	splitStr := strings.Split(s, "")
	fmt.Println(splitStr)
	lastIndex := len(splitStr) - 1

	for i := 0; i < len(splitStr)/2; i++ {
		if splitStr[i] != splitStr[lastIndex-i] {
			return false
		}
	}

	return true
}
