package main

import "fmt"

func sum_to_n_a(n int) int {
	// This solution has a time complexity of O(n) and a space complexity of O(n)
	if (n <= 0) {
		return 0
	} else {
		return n + sum_to_n_a(n-1)
	}
}

func sum_to_n_b(n int) int {
	// This solution has a time complexity of O(1) and a space complexity of O(1)
	return n * (n + 1) / 2
}

func sum_to_n_c(n int) int {
	// This solution has a time complexity of O(n) and a space complexity of O(1)
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(sum_to_n_a(5)) // 15
	fmt.Println(sum_to_n_b(5)) // 15
	fmt.Println(sum_to_n_c(5)) // 15
}