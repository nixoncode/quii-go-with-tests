// Package _4_arrays_and_slices
// Created by GoLand.
// User: nixon
// Date: 12/1/2024
// Time: 12:52
package _4_arrays_and_slices

func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
