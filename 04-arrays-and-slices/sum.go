// Package _4_arrays_and_slices
// Created by GoLand.
// User: nixon
// Date: 12/1/2024
// Time: 12:52
package _4_arrays_and_slices

func Sum(numbers [5]int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
