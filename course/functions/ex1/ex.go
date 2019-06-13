package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 5, 6, 7, 8, 9))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(10,6))

}

func sum(nums ...int) int {
	fmt.Printf("nums is %T\n", nums)
	res := 0
	for _, num := range nums {
		res += num
	}
	return res

	sort.Float64s([]float64{34})
}
