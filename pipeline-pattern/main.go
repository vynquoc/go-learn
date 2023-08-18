package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	first := firstLine(nums)
	second := secondLine(first)
	for item := range second {
		fmt.Println("RECEIVE: ", item)
	}
}

func firstLine(nums []int) chan int {
	result := make(chan int)

	go func() {
		for i := 0; i < len(nums); i++ {
			result <- nums[i]
		}
		close(result)
	}()

	return result
}

func secondLine(first chan int) chan int {
	result := make(chan int)

	go func() {
		for item := range first {
			result <- item * 2
		}
		close(result)
	}()

	return result
}
