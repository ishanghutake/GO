package main

import (
	"fmt"
	"time"
)

// func main() {
// 	started := time.Now()
// 	foods := []string{"mashroom pizza", "pasta", "kebab", "cake"}
// 	for _, f := range foods {
// 		cook(f)
// 	}
// 	fmt.Printf("done in %v\n", time.Since(started))
// }

// func cook(food string) {
// 	fmt.Printf("cooking %s...\n", food)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("done cooking %s\n", food)
// 	fmt.Println("")
// }
//////////////////////////////////////////////////////////////////////////////////////////////
// func main() {
// 	started := time.Now()
// 	foods := []string{"mashroom pizza", "pasta", "kebab", "cake"}
// 	var wg sync.WaitGroup
// 	wg.Add(len(foods))
// 	for _, f := range foods {
// 		go func(f string) {
// 			cook(f)
// 			wg.Done()
// 		}(f)
// 	}
// 	wg.Wait()
// 	fmt.Printf("done in %v\n", time.Since(started))
// }

// func cook(food string) {
// 	fmt.Printf("cooking %s...\n", food)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("done cooking %s\n", food)
// 	fmt.Println("")
// }

func main() {
	started := time.Now()
	foods := []string{"mashroom pizza", "pasta", "kabeb", "cake"}
	results := make(chan bool)
	for _, f := range foods {
		go func(f string) {
			cook(f)
			results <- true
		}(f)
	}
	for i := 0; i < len(foods); i++ {
		<-results
	}
	fmt.Printf("done in %v\n", time.Since(started))
}
func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
	fmt.Println("")
}
