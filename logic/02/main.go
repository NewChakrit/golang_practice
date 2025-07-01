package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a, b := swap("Hi", "Hello")

	fmt.Printf("a:%v, b:%v\n", a, b)
}
