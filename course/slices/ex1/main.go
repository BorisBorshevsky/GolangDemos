package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr [3]int
	fmt.Printf("%v - %T", arr, arr)
	arrChange(arr, 1, 100)
	fmt.Println()
	fmt.Printf("%v - %T", arr, arr)
	fmt.Println()

	var s1 []int
	fmt.Printf("%v - %T", s1, s1)
	fmt.Println()
	fmt.Printf("is s1 is a nil? %v", s1 == nil)
	fmt.Println()

	fmt.Printf("len %v, cap %v", len(s1), cap(s1))

	s1 = arr[0:2]
	fmt.Println()
	fmt.Printf("%v - %T", s1, s1)
	fmt.Println()
	fmt.Printf("len %v, cap %v", len(s1), cap(s1))

	s1 = append(s1, 3)
	fmt.Println()
	fmt.Printf("len %v, cap %v", len(s1), cap(s1))
	fmt.Println()
	fmt.Printf("%v - %T, %v - %T, ", s1, s1, arr, arr)

	s2 := make([]int, 0, 5)
	s2 = appendN(s2, 3, 2)
	fmt.Println(s2)

	appendSizes()

	s3 := make([]int, 2)
	fmt.Println("is s3 is a nil?", s3 == nil)

	//copy(s3, s1)
	//fmt.Println("s3", s3)
	//fmt.Println("s1", s1)

	//var s4 []int

	fmt.Println("SDASD", concat([]int{1, 3, 4}, []int{2}))

	fmt.Println(median([]float64{1, 2, 3}))
	fmt.Println(median([]float64{1, 2, 3, 4}))

}

func appendN(vals []int, val, n int) []int {
	for i := 0; i < n; i++ {
		vals = append(vals, n)
	}
	return vals

}

// array is passed by value!!
func arrChange(arr [3]int, i, val int) {
	arr[i] = val
}

// slice is passed by ref!!
func sliceChange(arr []int, i, val int) {
	arr[i] = val
}

func appendSizes() {
	currCap := 0
	var s []int
	for i := 0; i < 15000; i++ {
		s = append(s, i)
		if cap(s) != currCap {
			fmt.Println(currCap, "->", cap(s), cap(s)-currCap)
			currCap = cap(s)
		}
	}

}

func concat(a1, a2 []int) []int {
	return append(a1, a2...)
}

func median(values []float64) (float64, error) {
	sort.Float64s(values)

	if len(values) == 0 {
		return 0, fmt.Errorf("empty")
	}

	if len(values)%2 == 0 {
		return (values[len(values)/2 -1] + values[len(values)/2 ]) / 2, nil
	}

	return values[len(values)/2], nil
}
