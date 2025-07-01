package main

import (
	"fmt"
	"strings"
)

// reverse string

func main() {
	var hello = "hello"
	split := strings.Split(hello, "")
	fmt.Println(split)

	reverseSlice := make([]string, len(split))
	for i, j := len(split)-1, 0; j < len(split); i, j = i-1, j+1 {
		reverseSlice[j] = split[i]
	}

	fmt.Println(strings.Join(reverseSlice, ""))
}
