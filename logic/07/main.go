package main

import "fmt"

/* map to slice
สมมติมี map[string]int ให้ return slice ของ struct */

type Pair struct {
	Key   string
	Value int
}

func main() {
	val := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 7,
	}

	fmt.Println(MapToSlice(val))
}

func MapToSlice(m map[string]int) []Pair {
	result := []Pair{}
	for k, v := range m {
		result = append(result, Pair{k, v})
	}
	return result
}
