package main

import (
	"fmt"
	"strings"
)

// 01 : Reverse String
var hello string = "hello"

func main() {
	splitHello := strings.Split(hello, "")
	//fmt.Println(splitHello)

	reverseHello := make([]string, len(splitHello))
	for i, j := 0, len(splitHello)-1; i < len(splitHello); i, j = i+1, j-1 {
		reverseHello[i] = splitHello[j]

	}

	result := strings.Join(reverseHello, "")
	fmt.Println(result)
}
