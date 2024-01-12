// Package _4_arrays_and_slices
// Created by GoLand.
// User: nixon
// Date: 12/1/2024
// Time: 12:45
package _4_arrays_and_slices

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
