package main

import "fmt"

func main() {
	arr := []int{2, 2, 2, 3, 5, 7, 11, 13}
	// usingMap(arr)
	usingLoop(arr)
}

func usingMap(arr []int) {
	m := make(map[int]int)
	var va int
	for _, j := range arr {

		if _, ok := m[j]; ok {
			va = m[j]
			va = va + 1
			m[j] = va
		} else {
			m[j] = 1
		}
	}

	fmt.Print(m)
}

func usingLoop(arr []int) {
	fmt.Print("Using Loop")
}
